package services

import (
	"io"
	"log"

	"github.com/ebitengine/oto/v3"
	"github.com/pldcanfly/fylarm/internal/media"
)

type MediaService struct {
	Context *oto.Context
	Player  *oto.Player
}

type MediaStream interface {
	Stream() io.Reader
	Bitrate() int
	Artist() string
	Song() string
	Show() string
	Presenter() string
	Station() string
	Albumart() string
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
	}
}

func (ms *MediaService) PlayOE3() {
	if ms.Player != nil {
		ms.Player.Close()
	}

	// HERE BE GOFUNCS AND CONTEXTS!

	player := ms.Context.NewPlayer(media.NewOE3Stream().Stream())
	player.Play()

	for player.IsPlaying() {
	}
}
