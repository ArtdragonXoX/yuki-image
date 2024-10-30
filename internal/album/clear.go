package album

import (
	"fmt"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/utils"
)

func Clear(id uint64) error {
	album, err := db.SelectAlbum(id)
	if err != nil {
		return err
	}
	return ClearFromName(album.Name)
}

func ClearFromName(name string) error {
	id, err := SelectIdFromName(name)
	if err != nil {
		return err
	}
	err = db.ClearAlbum(id)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, name)
	err = utils.DeleteDir(path)
	_ = utils.EnsureDir(path)
	if err != nil {
		return err
	}
	return nil
}
