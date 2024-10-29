package format

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id uint64) (model.Format, error) {
	return db.SelectFormat(id)
}

func SelectIdFromName(name string) (uint64, error) {
	return db.SelectFromatIdFromName(name)
}

func SelectFormatFromName(name string) (model.Format, error) {
	return db.SelectFromatFromName(name)
}

func SelectAll() ([]model.Format, error) {
	return db.SelectAllFormat()
}
