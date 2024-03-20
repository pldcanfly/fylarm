package components

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewNext(next time.Time) *widget.Label {
	w := widget.NewLabel(fmt.Sprintf("Nächster Alarm: %s", next.Format("15:04")))
	w.Alignment = fyne.TextAlignCenter

	return w
}
