package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Component interface {
	Widget() *fyne.Widget
	Init() // Init(AlarmService)
}

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

// func getNext() *widget.Label {
// 	// n, err :=
// 	// This function needs to now the Alarmservice to execute .nextAlarm() on it:
// 	// Does it get it as a parameter?
// 	// Is there a central interface like a Server in typical http-stuff?
// 	// Is it strategy-patterned in?
//
// 	return l
// }

type NextAlarm struct {
	LabelWithStyle *widget.Label
}

func NewNext() *NextAlarm {

	w := widget.NewLabelWithStyle("Nächster Alarm: 13:37", LabelAlign, LabelStyle)
	return &NextAlarm{
		LabelWithStyle: w,
	}
}

func (w *NextAlarm) Init() {
	fmt.Println("Next Initialized")
}

func (w *NextAlarm) Widget() *widget.Label {
	return w.LabelWithStyle
}
