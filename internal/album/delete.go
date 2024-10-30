package album

import (
	"fmt"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/utils"
)

func Delete(id uint64) error {
	album, err := db.SelectAlbum(id)
	if err != nil {
		return err
	}
	err = db.DeleteAlbum(id)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, album.Name)
	err = utils.DeleteDir(path)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromName(name string) error {
	albumId, err := SelectIdFromName(name)
	if err != nil {
		return err
	}
	return Delete(albumId)
}
