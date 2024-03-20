package components

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

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
