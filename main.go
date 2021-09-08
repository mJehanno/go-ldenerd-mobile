package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	goldenerApp := app.New()
	mainWindow := goldenerApp.NewWindow("Goldener")
	purple := color.NRGBA{R: 136, G: 0, B: 136, A: 255}

	rect := canvas.NewRectangle(purple)
	mainWindow.SetContent(rect)

	mainWindow.Resize(fyne.NewSize(400, 400))

	/**  Login Form Dialog  */
	serverInput := widget.NewEntry()
	serverItem := widget.NewFormItem("Server", serverInput)

	loginInput := widget.NewEntry()
	loginItem := widget.NewFormItem("Login", loginInput)

	passInput := widget.NewPasswordEntry()
	passItem := widget.NewFormItem("Password", passInput)

	/** Gold Tab Container */
	goldContainer := container.NewVSplit(
		canvas.NewText("Current Coins : ", color.Black),
		canvas.NewText("Gold : ", color.Black),
	)

	/** History Tab Container */
	historyContainer := container.NewScroll(canvas.NewText("Gold : ", color.Black))

	/**  Tabs  */
	tabs := container.NewAppTabs(
		container.NewTabItem("Gold", goldContainer),
		container.NewTabItem("History", historyContainer),
	)

	loginDialog := dialog.NewForm("Login", "Login", "Cancel", []*widget.FormItem{serverItem, loginItem, passItem}, func(b bool) {
		if !b {
			panic("Can't do anything without login")
		}

		BaseUrl = serverInput.Text
		userLogin := loginInput.Text
		userPass := passInput.Text

		if err := Login(userLogin, userPass); err != nil {
			return
		}

		mainWindow.SetContent(tabs)

	}, mainWindow)

	loginDialog.Show()

	mainWindow.ShowAndRun()
}
