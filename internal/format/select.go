package format

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id uint64) (model.Format, error) {
	return db.SelectFormat(id)
}

func SelectAll() ([]model.Format, error) {
	return db.SelectAllFormat()
}
