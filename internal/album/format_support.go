package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func InsertFormatSupport(formatSupport model.FormatSupport) error {
	return db.InsertFormatSupport(formatSupport)
}
func DeleteFormatSupport(formatSupport model.FormatSupport) error {
	return db.DeleteFormatSupport(formatSupport)
}

func SelectFormatSupport(albumId uint64) ([]model.Format, error) {
	return db.SelectFormatSupport(albumId)
}
