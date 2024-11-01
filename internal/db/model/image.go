package model

import "time"

type Image struct {
	Key        string    `gorm:"key;primaryKey;not null;unique",json:"id"`
	Name       string    `gorm:"name;not null;unique",json:"name"`
	AlbumId    uint64    `gorm:"album_id;not null",json:"album_id"`
	Pathname   string    `gorm:"path_name;not null",json:"path_name"`
	OriginName string    `gorm:"origin_name;not null",json:"origin_name"`
	Size       uint64    `gorm:"size",json:"size"`
	Mimetype   string    `gorm:"mimetype" json:"mimetype"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
}
