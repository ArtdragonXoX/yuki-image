package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func SelectImage(id uint64, page uint64, size uint64) (model.ImageList, error) {
	return db.SelectImageFromAlbum(id, page, size)
}

func GetImageTotal(id uint64) (uint64, error) {
	return db.GetAlbumImageTotal(id)
}
