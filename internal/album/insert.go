package album

import (
	"fmt"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func Insert(album model.Album) (uint64, error) {
	id, err := db.InsertAlbum(album)
	if err != nil {
		return 0, err
	}
	pathname := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, album.Name)
	err = utils.EnsureDir(pathname)
	if err != nil {
		return 0, err
	}
	return id, nil
}
