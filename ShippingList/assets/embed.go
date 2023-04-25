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
