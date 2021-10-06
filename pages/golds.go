package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var Gold = ""
var GoldDetail = ""
var GoldBinding binding.String = binding.BindString(&Gold)
var GoldDetailBinding binding.String = binding.BindString(&GoldDetail)

func GetGoldPage(w fyne.Window) *fyne.Container {
	return container.NewGridWithRows(7,
		layout.NewSpacer(),
		container.NewVBox(
			container.NewCenter(widget.NewLabelWithData(GoldBinding)),
		),
		layout.NewSpacer(),
		container.NewVBox(widget.NewSeparator()),
		layout.NewSpacer(),
		container.NewVBox(
			container.NewCenter(widget.NewLabelWithData(GoldDetailBinding)),
		),
		layout.NewSpacer(),
	)

}
