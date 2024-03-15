package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/pldcanfly/fylarm/internal/components"
	"github.com/pldcanfly/fylarm/internal/components/alarm"
)

func main() {
	a := app.New()

	as := alarm.NewAlarmService()
	t, _ := time.Parse("15:04:05", "21:00:00")
	t2, _ := time.Parse("15:04:05", "22:00:00")
	t3, _ := time.Parse("15:04:05", "23:00:00")
	t4, _ := time.Parse("15:04:05", "11:00:00")
	t5, _ := time.Parse("15:04:05", "01:00:00")

	as.NewAlarm(t)
	as.NewAlarm(t2)
	as.NewAlarm(t3)
	as.NewAlarm(t4)
	as.NewAlarm(t5)

	w := a.NewWindow("fylarm")
	w.Resize(fyne.NewSize(800, 480))
	// w.SetFullScreen(true)
	w.SetPadded(false)
	w.SetContent(components.Layout())
	// w.ShowAndRun()
}
