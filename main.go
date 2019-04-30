package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/widgets"
)

var myapp *myApp

type myApp struct {
	app        *widgets.QApplication
	window     *widgets.QDialog
	grid       *widgets.QVBoxLayout
	board      *widgets.QTextEdit
	input      *widgets.QLineEdit
	popMenu    *widgets.QMenu
	itemStatus *widgets.QAction
	tray       *widgets.QSystemTrayIcon
	iconRun    *gui.QIcon
	iconStop   *gui.QIcon
	status     bool
	running    bool
	myCmd      *myCommand
	srv        *exec.Cmd
	chStaus    chan int
	pipeStdin  io.Writer
	once       sync.Once

	bridge *QtBridge
}

func (s *myApp) Run(name string) {
	s.app = widgets.NewQApplication(len(os.Args), os.Args)
	s.window = widgets.NewQDialog(nil, core.Qt__Window)
	s.app.SetActiveWindow(s.window)

	s.window.ConnectShowEvent(func(e *gui.QShowEvent) {
		s.once.Do(func() {
			home1, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			json1 := filepath.Join(home1, "config", name, "app.json")
			cfg, err := GetCfgFromJSON(json1)
			if err != nil {
				s.showHelp(json1)
				return
			}
			s.myCmd = new(myCommand)
			s.myCmd.Name = cfg.Exec
			args := []string{s.myCmd.Name}

			sp := strings.Split(cfg.Args, " ")
			for _, v := range sp {
				if len(v) == 0 {
					continue
				}
				args = append(args, v)
			}
			s.myCmd.Args = args
			//fmt.Println(s.myCmd.Args)

			sp = strings.Split(cfg.Envs, ";")
			envs := []string{}
			for _, v := range os.Environ() {
				envs = append(envs, v)
			}
			for _, v := range sp {
				if len(v) == 0 {
					continue
				}
				envs = append(envs, v)
			}
			s.myCmd.Envs = envs
			if len(cfg.Wd) > 0 {
				os.Chdir(cfg.Wd)
			}

			s.iconRun = gui.NewQIcon5(path.Join(home1, "config", name, "run.png"))
			s.iconStop = gui.NewQIcon5(path.Join(home1, "config", name, "stop.png"))
			s.window.SetWindowTitle("Controller-" + s.myCmd.Name)

			s.setSignals()

			s.setTray()
		})

	})

	s.window.SetFixedSize2(400, 300)

	s.grid = widgets.NewQVBoxLayout()

	s.board = widgets.NewQTextEdit(s.window)
	s.grid.AddWidget(s.board, 1, 0)
	s.board.SetTextColor(gui.NewQColor2(core.Qt__black))
	s.board.SetTextBackgroundColor(gui.NewQColor2(core.Qt__white))
	s.board.SetReadOnly(true)

	s.input = widgets.NewQLineEdit(s.window)
	s.grid.AddWidget(s.input, 1, 0)

	s.window.SetLayout(s.grid)
	s.window.ShowNormal()

	s.app.Exec()
}

func (s *myApp) setSignals() {
	//signals
	s.input.ConnectEditingFinished(func() {
		txt1 := s.input.Text()
		defer s.input.SetText("")
		s.pipeStdin.Write([]byte(txt1 + "\n"))
	})

	s.window.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		e.Ignore()
		s.window.Hide()
	})

	s.bridge = NewQtBridge(nil)

	s.bridge.ConnectUpdateRun(func() {
		//s.tray.Hide()
		s.tray.SetIcon(s.iconRun)
		s.tray.SetToolTip(s.myCmd.Name + " Running")
		s.itemStatus.SetText("Stop")
		s.status = true
		//s.tray.Show()
		log.Println("run " + s.myCmd.Name)
	})

	s.bridge.ConnectUpdateStop(func() {
		//s.tray.Hide()
		s.tray.SetIcon(s.iconStop)
		s.tray.SetToolTip(s.myCmd.Name + " Stopped")
		s.itemStatus.SetText("Run")
		s.status = false
		//s.tray.Show()
		s.board.Append(s.myCmd.Name + " : Killed\n")
		log.Println("kill " + s.myCmd.Name)
	})
}

func (s *myApp) setTray() {
	s.status = false
	s.chStaus = make(chan int, 1)
	go s.waitClose()

	s.popMenu = s.createMenu()

	s.tray = widgets.NewQSystemTrayIcon2(s.iconStop, s.window)
	s.tray.SetToolTip(s.myCmd.Name)
	s.tray.SetContextMenu(s.popMenu)
	s.tray.Show()

	s.changeStatus()
}

func (s *myApp) changeStatus() {
	if s.status {
		err := s.srv.Process.Kill()
		if err != nil {
			return
		}
		s.status = false
	} else {
		go s.runCmd()
		s.status = true
	}
}

func (s *myApp) createMenu() *widgets.QMenu {
	menu := widgets.NewQMenu(s.window)
	item := menu.AddAction("Quit")
	item.ConnectTriggered(func(checked bool) {
		if s.status {
			s.srv.Process.Kill()
		}
		s.window.DisconnectCloseEvent()
		s.window.Close()
	})

	s.itemStatus = menu.AddAction("Run")
	s.itemStatus.ConnectTriggered(func(checked bool) {
		s.changeStatus()
	})

	item = menu.AddAction("ShowWindow")
	item.ConnectTriggered(func(checked bool) {
		s.window.ShowNormal()
		s.window.Raise()
	})

	return menu
}

func (s *myApp) runCmd() {
	s.srv = s.myCmd.GetCmd()
	if s.srv == nil {
		s.chStaus <- 2
		return
	}
	r, _ := s.srv.StdoutPipe()
	go s.showOutput(r)
	r, _ = s.srv.StderrPipe()
	go s.showOutput(r)
	s.chStaus <- 1
	w, _ := s.srv.StdinPipe()
	defer w.Close()
	s.pipeStdin = w
	//fmt.Println("run")
	err := s.srv.Run()
	if err != nil {
		log.Println(err)
	}
	s.chStaus <- 3
	//fmt.Println("stop")
}

func (s *myApp) waitClose() {
	for {
		ret, ok := <-s.chStaus
		if !ok {
			s.window.Close()
		}
		if ret == 1 {
			s.signalRun()
			//fmt.Println("update run")
		} else if ret == 3 {
			s.signalStop()
			//fmt.Println("update stop")
		} else if ret == 2 {
			s.window.Close()
		}
	}
}

func (s *myApp) signalRun() {
	s.bridge.UpdateRun()
}

func (s *myApp) signalStop() {
	s.bridge.UpdateStop()
}

func (s *myApp) showOutput(r io.ReadCloser) {
	defer r.Close()
	for {
		var buf [480]byte
		n, err := r.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}

		s.board.Append(string(buf[:n]))

	}
}

func (s *myApp) showHelp(name string) {
	msg := `
	Usge: 
		tray-controller cfg-dir-name
	需要配置文件和2个图标！
	配置文件：` + name + ` 包含内容：
	{
		"exec":"/full/path/to/prog",
		"args":"-name2 value1 -name2 value2 ...",
		"envs":"Key1=Value1;Key2=Value2;...",
		"wd":"/path/to/work/dir"
	}
	"args"、"envs"、"wd"可以省略。
	图标和配置文件在同一目录，分别是：
	run.png ：代表正在运行
	stop.png ：代表停止状态`
	dlg := widgets.NewQMessageBox(s.window)
	dlg.SetWindowTitle("Need config file:" + name)
	dlg.SetToolTip("Need config file:" + name)
	dlg.SetText(msg)
	var once sync.Once
	dlg.ConnectChildEvent(func(e *core.QChildEvent) {
		once.Do(func() {
			s.window.Close()
			log.Println("Quit")
		})
	})
	dlg.SetModal(true)
	dlg.Show()
}

func main() {
	var name string
	if len(os.Args) == 2 {
		name = os.Args[1]
	}
	myapp = new(myApp)
	myapp.Run(name)
}
