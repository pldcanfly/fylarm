package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fylarm")

	clock := widget.NewLabel("")

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.SetContent(clock)
	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	t := time.Now().Format("Time: 15:04:05")
	clock.SetText(t)
}
