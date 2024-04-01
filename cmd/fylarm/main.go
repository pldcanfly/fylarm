package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/pldcanfly/fylarm/internal/services"
)

func main() {
	a := app.New()

	services.Alarm.NewAlarm(time.Now().Add(1 * time.Minute))
	w := a.NewWindow("fylarm")
	w.Resize(fyne.NewSize(800, 480))
	// w.SetFullScreen(true)
	w.SetPadded(false)
	l, err := services.Layout()
	if err != nil {
		panic(err)
	}
	w.SetContent(l)

	w.ShowAndRun()
}
