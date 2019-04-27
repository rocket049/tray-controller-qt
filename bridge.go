package main

import (
	"github.com/therecipe/qt/core"
)

type QtBridge struct {
	core.QObject

	_ func() `signal:"updateRun"`
	_ func() `signal:"updateStop"`
}
