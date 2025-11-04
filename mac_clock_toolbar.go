package main

import (
	"fmt"
	"github.com/caseymrm/menuet"
	"github.com/maxrichie5/mac-clock-toolbar/internal/cfg"
	"github.com/maxrichie5/mac-clock-toolbar/internal/clocks"
	"time"
)

func main() {
	cfg.Start()
	clocks.Start()
	go app()
	menuet.App().RunApplication()
	fmt.Println("Running")
}

func setMenuState() {
	menuet.App().Children = func() []menuet.MenuItem {
		menu := clocks.GetClocksMenu()
		menu = append(menu, menuet.MenuItem{Type: menuet.Separator})
		menu = append(menu, clocks.GetSettingsMenu()...)
		return menu
	}
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: clocks.GetActiveClocks(),
	})
	menuet.App().Label = "Mac Clock Toolbar"
}

func app() {
	for {
		setMenuState()
		time.Sleep(time.Second)
	}
}
