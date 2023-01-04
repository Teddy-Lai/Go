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

var searchBox = container.NewMax()
var check = false

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

	homeBox := container.NewMax(homeBtn, title)
	searchBox = search()
	tabContent := container.NewHBox(homeBox, allBtn, themeBtn, layout.NewSpacer(), searchBox)
	return tabContent
}

func search() *fyne.Container {
	searchEntry := widget.NewEntry()
	searchEntry.PlaceHolder = "想找什麼商品?"
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
	cancelBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		fmt.Println("取消")
		check = !check
		searchBox.Objects = []fyne.CanvasObject{search()}
	})
	searchIcon := tutorials.SearchIcon()
	searchBtn := widget.NewButtonWithIcon("", searchIcon, func() {
		fmt.Println("搜尋")
		check = !check
		searchBox.Objects = []fyne.CanvasObject{search()}
	})
	var content *fyne.Container
	if check {
		searchBtn.Importance = widget.HighImportance
		loginBtn.Hide()
		registerBtn.Hide()
		cartBtn.Hide()
		content = container.NewBorder(nil, nil, nil, container.NewHBox(cancelBtn, searchBtn), searchEntry)
	} else {
		searchBtn.Importance = widget.MediumImportance
		cancelBtn.Hide()
		searchEntry.Hide()
		content = container.NewHBox(layout.NewSpacer(), loginBtn, registerBtn, cartBtn, searchBtn)
	}
	bgRect := canvas.NewRectangle(color.White)
	bgRect.SetMinSize(fyne.NewSize(300, 40))
	contentBox := container.NewMax(bgRect, content)
	return contentBox
}
