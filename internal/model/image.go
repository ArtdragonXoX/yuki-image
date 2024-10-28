package model

const (
	JPEG uint64 = 1
	PNG  uint64 = 2
	GIF  uint64 = 3
)

type Image struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url,omitempty"`
	AlbumId    uint64 `json:"album_id,omitempty"`
	Pathname   string `json:"pathname"`
	OriginName string `json:"origin_name"`
	Size       uint64 `json:"size"`
	Mimetype   string `json:"mimetype"`
	Time       string `json:"time,omitempty"`
}
