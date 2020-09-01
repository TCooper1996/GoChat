package views

import (
	"github.com/jroimartin/gocui"
)

type MenuManager struct {
}

func (m MenuManager) Layout(g *gocui.Gui) error {
	maxx, maxy := g.Size()

	if v, err := g.SetView("Main", 0, 0, maxx, maxy); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		v.Title = "Main Menu"

	}
	return nil
}
