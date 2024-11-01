package album

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id uint64) (model.Album, error) {
	dbalbum, err := db.SelectAlbum(id)
	if err != nil {
		return model.Album{}, err
	}
	var album model.Album
	album.FromDBModel(dbalbum)
	return album, nil
}

func SelectAll() ([]model.Album, error) {
	dbalbums, err := db.SelectAllAlbum()
	if err != nil {
		return nil, err
	}
	var albums []model.Album
	for _, dbalbum := range dbalbums {
		var album model.Album
		album.FromDBModel(dbalbum)
		albums = append(albums, album)
	}
	return albums, nil
}

func SelectFromName(name string) (model.Album, error) {
	id, err := db.SelectAlbumIdFromName(name)
	if err != nil {
		return model.Album{}, err
	}
	return Select(id)
}

func SelectIdFromName(name string) (uint64, error) {
	return db.SelectAlbumIdFromName(name)
}
