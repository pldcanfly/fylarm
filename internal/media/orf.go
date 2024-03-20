package media

import (
	"io"
	"log"
	"net/http"

	"github.com/hajimehoshi/go-mp3"
)

type ORFStream struct {
	name    string
	metaurl string
	stream  io.Reader
}

// https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?items=true&_o=sound.orf.at
// https://orf-live.ors-shoutcast.at/oe3-q2a.m3u
func NewOE3Stream() *ORFStream {

	pl, err := http.Get("https://orf-live.ors-shoutcast.at/oe3-q2a.m3u")
	if err != nil {
		log.Printf("retrieving playlist: %v", err)
	}
	defer pl.Body.Close()

	bytes, err := io.ReadAll(pl.Body)
	if err != nil {
		log.Printf("reading playlist: %v", err)
	}

	resp, err := http.Get(string(bytes))
	if err != nil {
		log.Fatalf("opening stream: %v", err)
	}
	defer resp.Body.Close()

	dec, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		log.Fatalf("decoding stream: %v", err)
	}

	return &ORFStream{
		name:    "Hitradion Ö3",
		metaurl: "https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?_o=sound.orf.at",
		stream:  dec,
	}

}

func (s *ORFStream) Stream() io.Reader {
	return s.stream
}

func (s *ORFStream) Bitrate() int {
	return 48000
}

func (s *ORFStream) Artist() string {
	return "NYI-Artist"
}

func (s *ORFStream) Song() string {
	return "NYI-Song"
}

func (s *ORFStream) Show() string {
	return "NYI-Show"
}

func (s *ORFStream) Presenter() string {
	return "NYI-Kratky"
}

func (s *ORFStream) Station() string {
	return "NYI-Sender"
}

func (s *ORFStream) Albumart() string {
	return "NYI-Album"
}
