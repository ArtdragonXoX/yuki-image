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
	dbalbum, err := db.SelectAlbum(album.Id)
	if err != nil {
		return err
	}
	if album.MaxHeight == 0 {
		album.MaxHeight = dbalbum.MaxHeight
	}
	if album.MaxWidth == 0 {
		album.MaxWidth = dbalbum.MaxWidth
	}
	if album.Name == "" {
		album.Name = dbalbum.Name
	}
	return db.UpdateAlbum(album.ToDBModel())
}
