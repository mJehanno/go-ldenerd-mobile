package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var GoldBinding binding.String = binding.NewString()
var GoldDetailBinding binding.String = binding.NewString()

func GetGoldPage(w fyne.Window) *fyne.Container {
	gold, _ := GoldBinding.Get()
	goldDetail, _ := GoldDetailBinding.Get()

	return container.NewGridWithRows(7,
		layout.NewSpacer(),
		container.NewVBox(
			container.NewCenter(canvas.NewText("Gold Amount : "+gold, color.Black)),
		),
		layout.NewSpacer(),
		container.NewVBox(widget.NewSeparator()),
		layout.NewSpacer(),
		container.NewVBox(
			canvas.NewText("Coins Detail : "+goldDetail, color.Black),
		),
		layout.NewSpacer(),
	)

}
