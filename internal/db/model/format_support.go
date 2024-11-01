package model

type FormatSupport struct {
	FormatId uint64 `gorm:"column:format_id;primaryKey" json:"format_id"`
	AlbumId  uint64 `gorm:"column:album_id;primaryKey" json:"album_id"`
}
