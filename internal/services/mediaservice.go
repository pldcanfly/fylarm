package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/pldcanfly/fylarm/internal/media"
)

type MediaService struct {
	Context *oto.Context
	Player  *oto.Player
	Stream  media.MediaStream
	channel chan bool
}

var Media = NewMediaService()

func NewMediaService() *MediaService {
	ctx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   48000,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		log.Fatalf("oto context: %v", err)
	}
	<-ready

	return &MediaService{
		Context: ctx,
		channel: make(chan bool),
	}
}

func (ms *MediaService) Play(station string) (chan bool, error) {

	if ms.Player != nil {
		ms.Stop()
	}

	switch station {
	case "oe3":
		stream, rdy := media.NewOE3Stream()
		ms.PlayRemoteMP3(stream)
		return rdy, nil
	}
	return nil, fmt.Errorf("no such station")
}

func (ms *MediaService) Stop() {
	ms.channel <- true
}

func (ms *MediaService) PlayRemoteMP3(stream media.MediaStream) error {

	resp, err := http.Get(stream.Stream())

	if err != nil {
		return fmt.Errorf("opening stream: %v", err)
	}

	dec, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		return fmt.Errorf("decoding stream: %v", err)
	}
	ms.Player = ms.Context.NewPlayer(dec)
	ms.Stream = stream

	go func() {
		fmt.Println("play")
		ms.Player.Play()
		<-ms.channel
		ms.Player.Close()
		resp.Body.Close()
		fmt.Println("stopping")
	}()

	return nil
}
