package model

import (
	dbModel "yuki-image/internal/db/model"
)

const (
	JPEG uint64 = 1
	PNG  uint64 = 2
	GIF  uint64 = 3
)

type Image struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Url        string `json:"url,omitempty"`
	AlbumId    uint64 `json:"album_id,omitempty"`
	Pathname   string `json:"pathname"`
	OriginName string `json:"origin_name"`
	Size       uint64 `json:"size"`
	Mimetype   string `json:"mimetype"`
	Time       string `json:"time,omitempty"`
}

func (i *Image) ToDBModel() dbModel.Image {
	return dbModel.Image{
		Key:        i.Key,
		Name:       i.Name,
		AlbumId:    i.AlbumId,
		Pathname:   i.Pathname,
		OriginName: i.OriginName,
		Size:       i.Size,
		Mimetype:   i.Mimetype,
	}
}

func (i *Image) FromDBModel(dbImage dbModel.Image) {
	i.Key = dbImage.Key
	i.Name = dbImage.Name
	i.AlbumId = dbImage.AlbumId
	i.Pathname = dbImage.Pathname
	i.OriginName = dbImage.OriginName
	i.Size = dbImage.Size
	i.Mimetype = dbImage.Mimetype
	i.Time = dbImage.CreateTime.Format("2006-01-02 15:04:05")
}
