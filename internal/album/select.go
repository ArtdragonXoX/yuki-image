package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id uint64) (model.Album, error) {
	return db.SelectAlbum(id)
}

func SelectAll() ([]model.Album, error) {
	return db.SelectAllAlbum()
}
