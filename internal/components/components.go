package components

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

var LabelAlign = fyne.TextAlignCenter
var LabelStyle = fyne.TextStyle{}

func getNext() *widget.Label {
	l := widget.NewLabelWithStyle("Nächster Alarm: 20:00", LabelAlign, LabelStyle)
	// n, err :=
	// This function needs to now the Alarmservice to execute .nextAlarm() on it:
	// Does it get it as a parameter?
	// Is there a central interface like a Server in typical http-stuff?
	// Is it strategy-patterned in?

	return l
}
