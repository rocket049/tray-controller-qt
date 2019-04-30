package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/widgets"

	"github.com/rocket049/gettext-go/gettext"
)

type myCfg struct {
	Exec string
	Args string
	Envs string
	Wd   string
}

type myApp struct {
	app         *widgets.QApplication
	window      *widgets.QDialog
	button      *widgets.QPushButton
	cfgName     *widgets.QLineEdit
	cfgExec     *widgets.QLineEdit
	cfgArgs     *widgets.QLineEdit
	cfgEnvs     *widgets.QLineEdit
	cfgWd       *widgets.QLineEdit
	cfgRunIcon  *widgets.QLineEdit
	cfgStopIcon *widgets.QLineEdit
}

var app *myApp
var osID int

func getCfgDir(name string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	var dir1 string
	if len(name) > 0 {
		dir1 = filepath.Join(home, "config", name)
	} else {
		return "", err
	}

	os.MkdirAll(dir1, os.ModePerm)
	return dir1, nil
}

// func getBinPath(name string) (string, error) {
// 	if len(name) == 0 {
// 		return "", errors.New("Must give me a name.")
// 	}
// 	switch osID {
// 	case 0:
// 		home, err := os.UserHomeDir()
// 		if err != nil {
// 			return "", err
// 		}
// 		dir1 := filepath.Join(home, "bin")
// 		os.MkdirAll(dir1, os.ModePerm)
// 		return filepath.Join(dir1, name), nil
// 	case 1:
// 		exe, err := os.Executable()
// 		if err != nil {
// 			return "", err
// 		}
// 		dir1 := filepath.Dir(exe)
// 		return filepath.Join(dir1, name+".exe"), nil
// 	}
// 	return "", nil
// }

func getControllerPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir1 := filepath.Dir(exe)
	var res string
	switch osID {
	case 0:
		res = filepath.Join(dir1, "tray-controller-qt")
	case 1:
		res = filepath.Join(dir1, "tray-controller-qt.exe")
	}
	return res, nil
}

func getIconDir() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	dir1 := filepath.Dir(exe)
	return filepath.Join(dir1, "..", "share", "traycontroller", "icons")
}

func getLocaleDir() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	dir1 := filepath.Dir(exe)
	return filepath.Join(dir1, "..", "share", "traycontroller", "locale")
}

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *myApp) Run() {
	s.app = widgets.NewQApplication(len(os.Args), os.Args)
	s.window = widgets.NewQDialog(nil, core.Qt__Window)
	s.app.SetActiveWindow(s.window)
	grid := widgets.NewQGridLayout(s.window)

	label := widgets.NewQLabel2(gettext.T("Create New Controller Config File."), s.window, core.Qt__Widget)
	grid.AddWidget3(label, 0, 0, 1, 2, 0)

	label = widgets.NewQLabel2(gettext.T("ControllerName:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 1, 0, 0)

	s.cfgName = widgets.NewQLineEdit(s.window)
	s.cfgName.SetFixedWidth(300)
	grid.AddWidget(s.cfgName, 1, 1, 0)
	switch osID {
	case 0:
		s.cfgName.SetToolTip(gettext.T("I will generate Menu item") + ":ControllerName")
		s.cfgName.SetPlaceholderText(gettext.T("I will generate Menu item") + ":ControllerName")
	case 1:
		s.cfgName.SetToolTip(gettext.T("I will generate script") + ":launcher/ControllerName.vbs")
		s.cfgName.SetPlaceholderText(gettext.T("I will generate script") + ":launcher/ControllerName.vbs")
	}

	label = widgets.NewQLabel2(gettext.T("Exec:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 2, 0, 0)

	s.cfgExec = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgExec, 2, 1, 0)

	s.cfgExec.SetToolTip(gettext.T("DOUBLE Click to show Choose Dialog"))
	s.cfgExec.SetPlaceholderText(gettext.T("DOUBLE Click to show Choose Dialog"))

	label = widgets.NewQLabel2(gettext.T("Args:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 3, 0, 0)

	s.cfgArgs = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgArgs, 3, 1, 0)

	s.cfgArgs.SetToolTip(gettext.T("Split arguments with SPACE, example:") + "args1 arg2 ...")
	s.cfgArgs.SetPlaceholderText(gettext.T("Split arguments with SPACE, example:") + "args1 arg2 ...")

	label = widgets.NewQLabel2(gettext.T("Envs:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 4, 0, 0)

	s.cfgEnvs = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgEnvs, 4, 1, 0)

	s.cfgEnvs.SetToolTip(gettext.T("Split enviroment with ';', example:") + "Key1=Value1;Key2=Value2;...")
	s.cfgEnvs.SetPlaceholderText(gettext.T("Split enviroment with ';', example:") + "Key1=Value1;Key2=Value2;...")

	label = widgets.NewQLabel2(gettext.T("Wd:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 5, 0, 0)

	s.cfgWd = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgWd, 5, 1, 0)
	s.cfgWd.SetToolTip(gettext.T("DOUBLE Click to show Choose Dialog"))
	s.cfgWd.SetPlaceholderText(gettext.T("DOUBLE Click to show Choose Dialog"))

	label = widgets.NewQLabel2(gettext.T("RunIcon:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 6, 0, 0)

	s.cfgRunIcon = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgRunIcon, 6, 1, 0)
	s.cfgRunIcon.SetToolTip(gettext.T("DOUBLE Click to show Choose Dialog"))
	s.cfgRunIcon.SetPlaceholderText(gettext.T("DOUBLE Click to show Choose Dialog"))

	label = widgets.NewQLabel2(gettext.T("StopIcon:"), s.window, core.Qt__Widget)
	grid.AddWidget(label, 7, 0, 0)

	s.cfgStopIcon = widgets.NewQLineEdit(s.window)
	grid.AddWidget(s.cfgStopIcon, 7, 1, 0)
	s.cfgStopIcon.SetToolTip(gettext.T("DOUBLE Click to show Choose Dialog"))
	s.cfgStopIcon.SetPlaceholderText(gettext.T("DOUBLE Click to show Choose Dialog"))

	s.button = widgets.NewQPushButton2(gettext.T("OK"), s.window)
	grid.AddWidget3(s.button, 8, 0, 1, 2, 0)

	s.window.SetLayout(grid)
	s.window.Show()
	s.setSignals()
	s.app.Exec()
}

func (s *myApp) setSignals() {
	s.cfgExec.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {

		s.cfgExec.SetText(widgets.QFileDialog_GetOpenFileName(s.window, gettext.T("Choose a Executable"), ".", "*", "*", widgets.QFileDialog__ReadOnly))

	})

	s.cfgWd.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {

		s.cfgWd.SetText(widgets.QFileDialog_GetExistingDirectory(s.window, gettext.T("Choose a Dir"), ".", widgets.QFileDialog__ShowDirsOnly))

	})

	s.cfgRunIcon.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {

		s.cfgRunIcon.SetText(widgets.QFileDialog_GetOpenFileName(s.window, gettext.T("Choose a PNG Image"), "../share/traycontroller/icons", "PNG Files(*.png *.PNG)", "*", widgets.QFileDialog__ReadOnly))

	})

	s.cfgStopIcon.ConnectMouseDoubleClickEvent(func(e *gui.QMouseEvent) {

		s.cfgStopIcon.SetText(widgets.QFileDialog_GetOpenFileName(s.window, gettext.T("Choose a PNG Image"), "../share/traycontroller/icons", "PNG Files(*.png *.PNG)", "*", widgets.QFileDialog__ReadOnly))

	})

	s.button.ConnectClicked(func(checked bool) {
		s.makeConfig()
		s.window.Close()
	})
}

func copyFile(src, dst string, mode os.FileMode) error {
	fp1, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	defer fp1.Close()
	fp2, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fp2.Close()
	_, err = io.Copy(fp1, fp2)
	return err
}

func makeBinPath(progPath, name1 string) string {
	switch osID {
	case 0:
		return fmt.Sprintf("\"%s\" \"%s\"", progPath, name1)
	case 1:
		return fmt.Sprintf("\"\"\"%s\"\"\"+\" \"+\"\"\"%s\"\"\"", progPath, name1)
	}
	return fmt.Sprintf("\"%s\" \"%s\"", progPath, name1)
}

func makeLauncher(name1, binPath, iconPath string) error {
	if len(binPath) == 0 {
		return errors.New("Must provide program pathname")
	}
	switch osID {
	case 0:
		//linux
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		path1 := filepath.Join(home, ".local", "share", "applications", name1+".desktop")
		tmpl := `[Desktop Entry]
Encoding=UTF-8
Version=1.0
Type=Application
Terminal=false
Exec=` + binPath + `
Name=` + name1 + `
Icon=` + iconPath + `
Categories=GTK;Utility;
Comment=Tray Controller`
		ioutil.WriteFile(path1, []byte(tmpl), 0755)
	case 1:
		//windows
		exe1, _ := os.Executable()
		dir1 := filepath.Dir(exe1)
		dir2 := filepath.Join(dir1, "launcher")
		os.MkdirAll(dir2, os.ModePerm)
		path2 := filepath.Join(dir2, name1+".vbs")
		tmpl := `Set shell = Wscript.createobject("wscript.shell")

a = shell.run (` + binPath + `,0)
`
		ioutil.WriteFile(path2, []byte(tmpl), 0755)
	}
	return nil
}

func zeroPanic(s string) {
	name1 := strings.Trim(s, " /")
	if len(name1) == 0 {
		panic("get zero value.")
	}
}

func (s *myApp) makeConfig() {
	name1 := s.cfgName.Text()
	name1 = strings.Trim(name1, " ")
	zeroPanic(name1)

	//set program
	progPath, err := getControllerPath()
	errPanic(err)
	binPath := makeBinPath(progPath, name1)

	cfgDir, err := getCfgDir(name1)
	//copy icons
	runIconPath := filepath.Join(cfgDir, "run.png")
	src := s.cfgRunIcon.Text()
	zeroPanic(src)
	copyFile(src, runIconPath, 0644)

	stopIconPath := filepath.Join(cfgDir, "stop.png")
	src = s.cfgStopIcon.Text()
	zeroPanic(src)
	copyFile(src, stopIconPath, 0644)

	//create config files
	exec := s.cfgExec.Text()
	errPanic(err)
	zeroPanic(exec)

	args := s.cfgArgs.Text()

	envs := s.cfgEnvs.Text()

	wd := s.cfgWd.Text()

	cfg := &myCfg{Exec: exec, Args: args, Envs: envs, Wd: wd}
	buf, err := json.Marshal(cfg)
	errPanic(err)
	fp, err := os.Create(filepath.Join(cfgDir, "app.json"))
	errPanic(err)
	defer fp.Close()
	fp.Write(buf)

	//create launcher
	makeLauncher(name1, binPath, runIconPath)
}

func main() {

	gettext.BindTextdomain("config", getLocaleDir(), nil)
	gettext.Textdomain("config")
	app = new(myApp)
	app.Run()

	gettext.SaveLog()
}
