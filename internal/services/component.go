package services

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pldcanfly/fylarm/internal/components"
	"github.com/pldcanfly/fylarm/internal/media"
)

func Layout() (*fyne.Container, error) {

	next, err := Alarm.NextAlarm()
	if err != nil {
		return nil, fmt.Errorf("layout next %v:", err)
	}
	clock, date := components.GetClock()

	test := media.NewOE3Stream()
	test.Refresh()
	album, err := fyne.LoadResourceFromURLString(test.Albumart())
	if err != nil {
		return nil, fmt.Errorf("loading albumart: %v", err)
	}
	fmt.Println(album.Name())

	ac := canvas.NewImageFromResource(album)
	ac.SetMinSize(fyne.NewSize(200, 200))

	c := container.NewStack(
		canvas.NewRectangle(color.Black),
		container.NewGridWithRows(2,
			container.NewVBox(
				layout.NewSpacer(),
				widget.NewLabelWithStyle(test.Station(), fyne.TextAlignCenter, fyne.TextStyle{}),
				container.NewCenter(
					ac,
				),
				widget.NewLabelWithStyle(test.Show(), fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(fmt.Sprintf("%s - %s", test.Artist(), test.Song()), fyne.TextAlignCenter, fyne.TextStyle{}),
				clock,
				date,
			),
			container.NewVBox(
				layout.NewSpacer(),
				components.NewNext(next.NextRing),
				components.GetWeather(),
				widget.NewLabelWithStyle(
					"Heute in Graz",
					fyne.TextAlignCenter,
					fyne.TextStyle{},
				),
				layout.NewSpacer(),
			),
		),
	)

	return c, nil
}

// I don't like this pattern, doesn't feel very idiomatic
// Why is it HERE. Why do i need to return arbitrary. Works for now WILL be replaced with
//
//	something like a componentstore and interfaces for components that implement an
//	Init() function and can handle stuff like that on the fly.
//	GO is good with concurrency and should be able to handle 2 threads seperatly updating
//	clock and date, but maybe I can merge both into one, which would be the best approach

// func getNext() *widget.Label {
// 	// n, err :=
// 	// This function needs to now the Alarmservice to execute .nextAlarm() on it:
// 	// Does it get it as a parameter?
// 	// Is there a central interface like a Server in typical http-stuff?
// 	// Is it strategy-patterned in?
//
// 	return l
// }
