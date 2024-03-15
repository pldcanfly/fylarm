package components

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// I don't like this pattern, doesn't feel very idiomatic
// Why is it HERE. Why do i need to return arbitrary. Works for now WILL be replaced with
//
//	something like a componentstore and interfaces for components that implement an
//	Init() function and can handle stuff like that on the fly.
//	GO is good with concurrency and should be able to handle 2 threads seperatly updating
//	clock and date, but maybe I can merge both into one, which would be the best approach
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
