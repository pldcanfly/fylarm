package media

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hajimehoshi/go-mp3"
)

type ORFStream struct {
	name    string
	metaurl string
	stream  io.Reader
	meta    *ORFMetaResponse
}

type ORFMetaResponse struct {
	Payload struct {
		Show      string `json:"title"`
		Moderator string `json:"moderator"`
		Item      struct {
			Artist string `json:"interpreter"`
			Song   string `json:"title"`
			Images []struct {
				Versions []struct {
					Albumart string `json:"path"`
				}
			} `json:"images"`
		} `json:"item"`
	} `json:"payload"`
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

	resp, err := http.Get(strings.TrimSuffix(string(bytes), "\r\n"))

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

func (s *ORFStream) Refresh() error {
	url := "https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?items=true&_o=sound.orf.at"

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)

	var data *ORFMetaResponse
	err = json.Unmarshal(r, &data)
	if err != nil {
		return err
	}

	s.meta = data

	return nil
}

func (s *ORFStream) Bitrate() int {
	return 48000
}

func (s *ORFStream) Artist() string {
	return s.meta.Payload.Item.Artist
}

func (s *ORFStream) Song() string {
	return s.meta.Payload.Item.Song
}

func (s *ORFStream) Show() string {
	return s.meta.Payload.Show
}

func (s *ORFStream) Moderator() string {
	return s.meta.Payload.Moderator
}

func (s *ORFStream) Station() string {
	return "Hitradio Ö3"
}

func (s *ORFStream) Albumart() string {
	return s.meta.Payload.Item.Images[0].Versions[0].Albumart
}
