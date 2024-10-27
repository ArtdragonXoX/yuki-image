package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Update(album model.Album, id uint64) error {
	album.Id = id
	return db.UpdateAlbum(album)
}
