package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mjehanno/go-ldenerd-mobile/service"
)

func GetLoginScreen(w fyne.Window, newC *fyne.Container) *fyne.Container {
	/**  Login Form Dialog  */
	serverInput := widget.NewEntry()
	serverInput.PlaceHolder = "Server"
	loginInput := widget.NewEntry()
	loginInput.PlaceHolder = "Username"
	passInput := widget.NewPasswordEntry()
	passInput.PlaceHolder = "Password"
	errorContainer := container.NewVBox()

	loginCardContainer := container.NewVBox(
		errorContainer,
		serverInput, loginInput, passInput,
	)

	loginButton := widget.NewButton("Login", func() {
		service.BaseUrl.Set(serverInput.Text)
		login := loginInput.Text
		pass := passInput.Text
		if err := service.Login(login, pass); err != nil {
			if len(errorContainer.Objects) == 0 {
				errorContainer.Add(
					canvas.NewText("Wrong configuration (server  / login / password)", color.RGBA{150, 0, 0, 245}),
				)
			}
			errorContainer.Refresh()
		} else {
			GoldBinding.Set(service.GetGold())
			GoldDetailBinding.Set(service.GetGoldDetail())
			w.SetContent(newC)

		}
	})

	loginCardContainer.Add(loginButton)
	loginCard := widget.NewCard("Login", "", loginCardContainer)

	c := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			loginCard,
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	return c
}
