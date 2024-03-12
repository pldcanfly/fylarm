package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pldcanfly/fylarm/internal/components"
)

func clocklayout(clock, date *canvas.Text) *fyne.Container {
	return container.NewStack(
		canvas.NewRectangle(color.Black),
		container.NewGridWithRows(2,
			container.NewVBox(
				layout.NewSpacer(),
				clock,
				date,
			),
			container.NewVBox(
				widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle("Nächster Alarm: 12:34", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle("Sonnig 24° Regenwahrscheinlichkeit: 50%", fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle("Heute in Graz", fyne.TextAlignCenter, fyne.TextStyle{}),
			),
		),
	)
}

func main() {
	a := app.New()

	w := a.NewWindow("fylarm")
	w.Resize(fyne.NewSize(1920, 1080))
	// w.SetFullScreen(true)
	w.SetPadded(false)
	w.SetContent(clocklayout(components.InitComponents()))
	w.ShowAndRun()
}
