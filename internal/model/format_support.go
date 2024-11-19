package model

import (
	"yuki-image/internal/db/model"
	dbModel "yuki-image/internal/db/model"
)

type FormatSupport struct {
	FormatId   uint64 `json:"format_id",omitempty`
	FormatName string `json:"format_name",omitempty`
	AlbumId    uint64 `json:"album_id",omitempty`
	AlbumName  string `json:"album_name",omitempty`
}

func (f *FormatSupport) ToDBModel() dbModel.FormatSupport {
	return dbModel.FormatSupport{
		FormatId: f.FormatId,
		AlbumId:  f.AlbumId,
	}
}

func (f *FormatSupport) FromDBModel(dbModel dbModel.FormatSupport) {
	f.FormatId = dbModel.FormatId
	f.AlbumId = dbModel.AlbumId
}

func ContainsFormatSupport(formatSupports []model.Format, format uint64) bool {
	for _, v := range formatSupports {
		if v.Id == format {
			return true
		}
	}
	return false
}
