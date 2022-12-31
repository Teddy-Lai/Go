package main

import (
	"ElectronicBusinessPlatform/tutorials"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("電商UI")

	title := canvas.NewText("名稱", color.White)
	title.TextSize = 30
	title.TextStyle.Bold = true
	btn := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		fmt.Println("Button Click!")
	})

	content := container.NewVBox(title, btn)
	// set theme
	t := &tutorials.MyTheme{}
	a.Settings().SetTheme(t)

	w.SetContent(content)
	w.ShowAndRun()
}
