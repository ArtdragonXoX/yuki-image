package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Update(album model.Album) error {
	var err error
	if album.Id == 0 {
		album.Id, err = SelectIdFromName(album.Name)
		if err != nil {
			return err
		}
	}
	return db.UpdateAlbum(album)
}
