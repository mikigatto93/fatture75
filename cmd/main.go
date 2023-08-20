package main

import (
	"fatture75/controller"
	"fatture75/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	controller.SetupController()

	a := app.New()

	//light theme is deprecated!
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("Fatture75")

	w.SetContent(view.GetMainWindowView(w))
	w.Resize(fyne.NewSize(700, 400))

	w.ShowAndRun()

}
