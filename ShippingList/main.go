package main

import (
	"ShippingList/tutorials"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("出貨單")
	w.Resize(fyne.NewSize(1600, 900))
	
	// set theme
	t := &tutorials.MyTheme{}
	a.Settings().SetTheme(t)

	w.SetContent(container.NewVBox())
	w.ShowAndRun()
}
