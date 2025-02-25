package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

var menu *fyne.Menu
var caffeinateMenuItem *fyne.MenuItem
var caffeinateCmd *exec.Cmd

func getCaffeinateTitle(isEnabled bool) string {
	return fmt.Sprintf("Caffeinate enabled: %t", isEnabled)
}

func toggleCaffeinate(isEnabled *bool) {
	if *isEnabled {
		caffeinateCmd = exec.Command("caffeinate", "-d")
		err := caffeinateCmd.Start()
		if err != nil {
			fmt.Println("Caffeinate başlatılamadı:", err)
			*isEnabled = false
		}
	} else {
		// Caffeinate işlemini durdur
		if caffeinateCmd != nil {
			err := caffeinateCmd.Process.Kill()
			if err != nil {
				fmt.Println("Caffeinate durdurulamadı:", err)
			}
			caffeinateCmd = nil
		}
	}
}

func main() {
	isCaffeinateEnabled := false
	a := app.New()
	w := a.NewWindow("Hello World")

	if desktop, ok := a.(desktop.App); ok {

		caffeinateMenuItem = fyne.NewMenuItem(getCaffeinateTitle(isCaffeinateEnabled), func() {
			isCaffeinateEnabled = !isCaffeinateEnabled
			toggleCaffeinate(&isCaffeinateEnabled)

			caffeinateMenuItem.Label = getCaffeinateTitle(isCaffeinateEnabled)
			menu.Refresh()
		})

		menu = fyne.NewMenu("Ctlkt", caffeinateMenuItem)
		desktop.SetSystemTrayMenu(menu)
	}

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
