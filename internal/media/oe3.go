package media

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type OE3Stream struct {
	name    string
	metaurl string
	stream  string
	meta    *OE3MetaResponse
	rdy     chan bool
}

type OE3MetaResponse struct {
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
			End string `json:"end"`
		} `json:"item"`
	} `json:"payload"`
}

// https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?items=true&_o=sound.orf.at
// https://orf-live.ors-shoutcast.at/oe3-q2a.m3u
func NewOE3Stream() (*OE3Stream, chan bool) {

	pl, err := http.Get("https://orf-live.ors-shoutcast.at/oe3-q2a.m3u")
	if err != nil {
		log.Printf("retrieving playlist: %v", err)
	}
	defer pl.Body.Close()

	bytes, err := io.ReadAll(pl.Body)
	if err != nil {
		log.Printf("reading playlist: %v", err)
	}

	c := make(chan bool)
	s := &OE3Stream{
		name:    "Hitradion Ö3",
		metaurl: "https://audioapi.orf.at/oe3/api/json/5.0/broadcast/onair?_o=sound.orf.at",
		stream:  strings.TrimSuffix(string(bytes), "\r\n"),
		rdy:     c,
	}
	go s.Refresh()

	return s, c

}

func (s *OE3Stream) Refresh() {
	for {
		fmt.Println("refresh")
		url := s.metaurl
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}
		defer res.Body.Close()

		r, err := io.ReadAll(res.Body)

		var data *OE3MetaResponse
		err = json.Unmarshal(r, &data)
		if err != nil {
			log.Println(err)
			continue
		}

		s.meta = data

		end, err := time.Parse(time.RFC3339Nano, s.meta.Payload.Item.End)
		if err != nil || time.Until(end) <= 0 {
			log.Println(err, "waiting 10s")
			time.Sleep(10 * time.Second)
			continue
		}
		end.Add(5 * time.Second)

		s.rdy <- true

		log.Println("Next Song: ", time.Until(end))
		time.Sleep(time.Until(end))
	}
}

func (s *OE3Stream) Type() int {
	return RemoteTypeMP3
}

func (s *OE3Stream) Stream() string {
	return s.stream
}

func (s *OE3Stream) Bitrate() int {
	return 48000
}

func (s *OE3Stream) Artist() string {
	return s.meta.Payload.Item.Artist
}

func (s *OE3Stream) Song() string {
	return s.meta.Payload.Item.Song
}

func (s *OE3Stream) Show() string {
	return s.meta.Payload.Show
}

func (s *OE3Stream) Moderator() string {
	return s.meta.Payload.Moderator
}

func (s *OE3Stream) Station() string {
	return "Hitradio Ö3"
}

func (s *OE3Stream) Albumart() string {
	if len(s.meta.Payload.Item.Images) >= 1 && len(s.meta.Payload.Item.Images[0].Versions) >= 1 {
		return s.meta.Payload.Item.Images[0].Versions[0].Albumart
	}
	return ""
}
