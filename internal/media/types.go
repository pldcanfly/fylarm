package media

type MediaStream interface {
	Refresh()
	Stream() string
	Bitrate() int
	Artist() string
	Song() string
	Show() string
	Moderator() string
	Station() string
	Albumart() (string, error)
}

const (
	RemoteTypeMP3 = iota
)
