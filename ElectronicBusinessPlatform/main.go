package main

import (
	"ElectronicBusinessPlatform/tutorials"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("電商UI")
	w.Resize(fyne.NewSize(1600, 900))

	separator := widget.NewSeparator()

	content := container.NewVBox(makeTabBtns(), separator)
	// set theme
	t := &tutorials.MyTheme{}
	a.Settings().SetTheme(t)

	w.SetContent(content)
	w.ShowAndRun()
}

func makeTabBtns() *fyne.Container {
	title := canvas.NewText("快電商", color.Black)
	title.TextSize = 30
	title.TextStyle.Bold = true
	homeBtn := widget.NewButton("", func() {
		fmt.Println("首頁")
	})
	allBtn := widget.NewButton("全部分類", func() {
		fmt.Println("全部分類")
	})
	themeBtn := widget.NewButton("主題", func() {
		fmt.Println("主題")
	})
	loginBtn := widget.NewButton("登入", func() {
		fmt.Println("登入")
	})
	registerBtn := widget.NewButton("註冊", func() {
		fmt.Println("註冊")
	})
	cartIcon := tutorials.CartIcon()
	cartBtn := widget.NewButtonWithIcon("", cartIcon, func() {
		fmt.Println("購物車")
	})
	searchIcon := tutorials.SearchIcon()
	searchBtn := widget.NewButtonWithIcon("", searchIcon, func() {
		fmt.Println("搜尋")
	})

	searchEntry := widget.NewEntry()
	searchEntry.PlaceHolder = "想找什麼商品?"
	cancelBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		fmt.Println("取消")
	})


	homeBox := container.NewMax(homeBtn, title)
	searchBox := container.NewHBox(searchEntry, loginBtn, registerBtn, cartBtn, cancelBtn, searchBtn)
	tabContent := container.NewHBox(homeBox, allBtn, themeBtn, layout.NewSpacer(), searchBox)
	return tabContent
}
