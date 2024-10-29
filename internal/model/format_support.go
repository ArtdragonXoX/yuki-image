package model

type FormatSupport struct {
	FormatId   uint64 `json:"format_id",omitempty`
	FormatName string `json:"format_name",omitempty`
	AlbumId    uint64 `json:"album_id",omitempty`
	AlbumName  string `json:"album_name",omitempty`
}
