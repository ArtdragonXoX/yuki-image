package model

type Image struct {
	Id         string  `json:"key"`
	Name       string  `json:"name"`
	Url        string  `json:"url"`
	Album      string  `json:"album"`
	Pathname   string  `json:"pathname"`
	OriginName string  `json:"origin_name"`
	Size       float64 `json:"size"`
	Mimetype   string  `json:"mimetype"`
	Time       string  `json:"time"`
}
