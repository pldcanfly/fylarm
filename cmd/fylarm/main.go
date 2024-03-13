package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/pldcanfly/fylarm/internal/components"
)

func main() {
	a := app.New()

	w := a.NewWindow("fylarm")
	w.Resize(fyne.NewSize(1920, 1080))
	// w.SetFullScreen(true)
	w.SetPadded(false)
	w.SetContent(components.Layout())
	w.ShowAndRun()
}
