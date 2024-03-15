package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Layout() *fyne.Container {
	clock, date := initComponents()
	return container.NewStack(
		canvas.NewRectangle(color.Black),
		container.NewGridWithRows(2,
			container.NewVBox(
				layout.NewSpacer(),
				clock,
				date,
			),
			container.NewVBox(
				widget.NewLabelWithStyle(
					"",
					fyne.TextAlignCenter,
					fyne.TextStyle{},
				),

				widget.NewLabelWithStyle(
					"Nächster Alarm: 12:34",
					fyne.TextAlignCenter,
					fyne.TextStyle{},
				),
				widget.NewLabelWithStyle(
					"Sonnig 24° Regenwahrscheinlichkeit: 50%",
					fyne.TextAlignCenter,
					fyne.TextStyle{},
				),
				widget.NewLabelWithStyle(
					"Heute in Graz",
					fyne.TextAlignCenter,
					fyne.TextStyle{},
				),
			),
		),
	)
}
