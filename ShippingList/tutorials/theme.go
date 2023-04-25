package tutorials

import (
	"ShippingList/assets"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// MyTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type MyTheme struct {
}

func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Bold {
		return assets.ResourceMSJHBD
	}
	return assets.ResourceMSJH
}

func (m MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
