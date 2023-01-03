package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed fonts/MSJH.ttf
var FontsMSJH []byte
var ResourceMSJH = &fyne.StaticResource{
	StaticName:    "MSJH.ttf",
	StaticContent: FontsMSJH,
}

//go:embed fonts/MSJHBD.ttf
var FontsMSJHBD []byte
var ResourceMSJHBD = &fyne.StaticResource{
	StaticName:    "MSJHBD.ttf",
	StaticContent: FontsMSJHBD,
}

//go:embed icons/cart.svg
var IconsCart []byte
var ResourceCart = &fyne.StaticResource{
	StaticName:    "cart.svg",
	StaticContent: IconsCart,
}

//go:embed icons/search.svg
var IconsSearch []byte
var ResourceSearch = &fyne.StaticResource{
	StaticName:    "search.svg",
	StaticContent: IconsSearch,
}
