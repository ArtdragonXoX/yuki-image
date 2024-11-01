package model

import dbModel "yuki-image/internal/db/model"

type Album struct {
	Id            uint64    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	MaxHeight     uint64    `json:"max_height,omitempty"`
	MaxWidth      uint64    `json:"max_width,omitempty"`
	FormatSupport []Format  `json:"format_support,omitempty"`
	UpdateTime    string    `json:"update_time,omitempty"`
	CreateTime    string    `json:"create_time,omitempty"`
	Image         ImageList `json:"image,omitempty"`
}

type ImageList struct {
	Image []Image `json:"image,omitempty"`
	Page  uint64  `json:"page,omitempty"`
	Size  uint64  `json:"size,omitempty"`
	Total uint64  `json:"total"`
}

func (a *Album) ToDBModel() dbModel.Album {
	return dbModel.Album{
		Id:        a.Id,
		Name:      a.Name,
		MaxHeight: a.MaxHeight,
		MaxWidth:  a.MaxWidth,
	}
}

func (a *Album) FromDBModel(album dbModel.Album) {
	a.Id = album.Id
	a.Name = album.Name
	a.MaxHeight = album.MaxHeight
	a.MaxWidth = album.MaxWidth
	a.UpdateTime = album.UpdateTime.Format("2006-01-02 15:04:05")
	a.CreateTime = album.CreateTime.Format("2006-01-02 15:04:05")
}
