package services

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/pldcanfly/fylarm/internal/components"
)

type Component interface {
	Widget() *fyne.Widget
	Init() // Init(AlarmService)
}

type ComponentService struct {
	Components    []*Component
	TopComponents []*Component
}

func NewComponentService() *ComponentService {
	return &ComponentService{}
}

func (cc *ComponentService) Register(c *Component) {
	cc.Components = append(cc.Components, c)
}

func (cc *ComponentService) Init() {
	for _, comp := range cc.Components {
		(*comp).Init()
	}
}

func Layout() *fyne.Container {

	c := container.NewVBox()
	next, err := Alarm.NextAlarm()
	if err != nil {
		log.Printf("layout next %v:", err)
	} else {

		c.Add(components.NewNext(next.NextRing))
	}
	return c
	// clock, date := initComponents()
	// return container.NewStack(
	// 	canvas.NewRectangle(color.Black),
	// 	container.NewGridWithRows(2,
	// 		container.NewVBox(
	// 			layout.NewSpacer(),
	// 			clock,
	// 			date,
	// 		),
	// 		container.NewVBox(
	// 			layout.NewSpacer(),
	// 			NewNext().Widget(),
	// 			widget.NewLabelWithStyle(
	// 				"Sonnig 24° Regenwahrscheinlichkeit: 50%",
	// 				fyne.TextAlignCenter,
	// 				fyne.TextStyle{},
	// 			),
	// 			widget.NewLabelWithStyle(
	// 				"Heute in Graz",
	// 				fyne.TextAlignCenter,
	// 				fyne.TextStyle{},
	// 			),
	// 			layout.NewSpacer(),
	// 		),
	// 	),
	// )
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
