package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/mjehanno/go-ldenerd-mobile/pages"
)

var IsUserLogged bool = false

func main() {
	goldenerApp := app.New()
	mainWindow := goldenerApp.NewWindow("Goldener")
	mainWindow.Resize(fyne.NewSize(800, 700))
	goldenerApp.SetIcon(resourceGobelinJpg)
	gTab := container.NewTabItem("Golds", pages.GetGoldPage(mainWindow))
	tTab := container.NewTabItem("Transactions", pages.GetTransactionsPage(mainWindow))
	tabs := container.NewAppTabs(gTab, tTab)
	home := container.NewMax(tabs)
	loginPage := pages.GetLoginScreen(mainWindow, home)
	mainWindow.CenterOnScreen()
	mainWindow.SetContent(loginPage)
	mainWindow.ShowAndRun()
}

func handlepanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}
