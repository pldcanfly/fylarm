package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/pldcanfly/fylarm/internal/alarm"
	"github.com/pldcanfly/fylarm/internal/ui"
)

func main() {
	a := app.New()
	// https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?items=true&_o=sound.orf.at
	// Hier kann man Meta-Daten rausziehen
	// Maybe ein Interface dafür machen und Stream als "Stream" behandeln und fertig?

	playlist, err := http.Get("https://orf-live.ors-shoutcast.at/oe3-q2a.m3u")
	if err != nil {
		log.Fatal("Couldn't load playlist", err)
	}
	defer playlist.Body.Close()
	bytes, _ := io.ReadAll(playlist.Body)
	fmt.Println(string(bytes), playlist.Header)

	context, rch, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		log.Fatalf("Failed to open speaker: %v", err)
	}

	<-rch

	resp, err := http.Get("http://ors-sn06.ors-shoutcast.at/oe3-q2a")
	if err != nil {
		log.Fatalf("HTTP Get: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)

	dec, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		log.Fatalf("Deocoder: %v", err)
	}

	player := context.NewPlayer(dec)
	player.Play()

	for player.IsPlaying() {
	}

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
	w.SetContent(ui.Layout())
	// w.ShowAndRun()
}
