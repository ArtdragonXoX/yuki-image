package model

import "time"

type Album struct {
	Id            uint64    `gorm:"id;primaryKey;autoIncrement:true" json:"id"`
	Name          string    `gorm:"name;not null;unique" json:"name"`
	MaxHeight     uint64    `gorm:"max_height;not null" json:"max_height"`
	MaxWidth      uint64    `gorm:"max_width;not null" json:"max_width"`
	UpdateTime    time.Time `gorm:"update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"create_time" json:"create_time"`
	Image         []*Image  `gorm:"foreignKey:AlbumId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FormatSupport []*Format `gorm:"many2many:format_supports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;association_jointable_foreignkey:format_id;jointable_foreignkey:album_id"`
}
