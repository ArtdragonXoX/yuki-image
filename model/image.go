package model

type Image struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Url        string  `json:"url"`
	AlbumId    uint64  `json:"album_id"`
	Pathname   string  `json:"pathname"`
	OriginName string  `json:"origin_name"`
	Size       float64 `json:"size"`
	Mimetype   string  `json:"mimetype"`
	Time       string  `json:"time"`
}
