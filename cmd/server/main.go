package main

import (
	Views "github.com/GoChat/"
	M "../../internal/server/Models"
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func main(){
	M.InitRoom()
	M.InitUser()
	server := M.NewUser("Server", nil)
	M.NewRoom("Main", server.ID)


	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Printf("Handle this")
		panic(err.Error())
	}

	g.SetManager(Views.MenuManager{})

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil{
		log.Panicln(err)
	}


	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit{

	}

}

func quit(g *gocui.Gui, v *gocui.View) error{
	return gocui.ErrQuit
}