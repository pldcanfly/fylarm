package components

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func initComponents() (*canvas.Text, *canvas.Text) {
	c := getClock()
	d := getDate()

	updateTime(c, d)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(c, d)
		}
	}()

	return c, d
}

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

func getClock() *canvas.Text {
	clock := canvas.NewText("", color.White)
	clock.TextSize = 50
	clock.Alignment = fyne.TextAlignCenter

	return clock
}

func getDate() *canvas.Text {
	date := canvas.NewText("", color.White)
	date.TextSize = 20
	date.Alignment = fyne.TextAlignCenter

	return date
}

func updateTime(c *canvas.Text, d *canvas.Text) {
	t := time.Now().Format("15:04:05")
	dt := time.Now().Format("Monday, 2 January")
	d.Text = dt
	c.Text = t
	d.Refresh()
	c.Refresh()
}
