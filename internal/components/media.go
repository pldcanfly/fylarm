package components

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pldcanfly/fylarm/internal/media"
)

func GetMedia(m media.MediaStream, rdy chan bool) *fyne.Container {
	cover := container.NewCenter()
	station := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{})
	song := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{})

	go updateText(m, station, song, cover, rdy)

	return container.NewVBox(
		station,
		cover,
		song,
	)

}

func updateText(m media.MediaStream, station *widget.Label, song *widget.Label, cover *fyne.Container, rdy chan bool) {
	for {
		<-rdy
		if m.Albumart() != "" {
			art, err := fyne.LoadResourceFromURLString(m.Albumart())
			if err != nil {
				log.Fatal("loading album art")
				return
			}
			ac := canvas.NewImageFromResource(art)
			ac.SetMinSize(fyne.NewSize(200, 200))
			cover.RemoveAll()
			cover.Add(ac)
		}

		station.SetText(fmt.Sprintf("%s: %s", m.Station(), m.Show()))
		if m.Artist() != "" {
			song.SetText(fmt.Sprintf("%s - %s", m.Artist(), m.Song()))
		} else {
			song.SetText(m.Song())
		}
	}
}
