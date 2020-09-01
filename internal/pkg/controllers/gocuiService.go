package controllers

import (
	V "GoChat/internal/pkg/views"
	"github.com/jroimartin/gocui"
	"log"
)

func InitGui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err.Error())
	}

	g.SetManager(V.MenuManager{})

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {

	}

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
