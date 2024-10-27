package image

import (
	"fmt"
	"os"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
)

func Delete(id string) error {
	image, err := db.SelectImage(id)
	if err != nil {
		return err
	}
	err = db.DeleteImage(id)
	if err != nil {
		return err
	}
	pathname := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, image.Pathname)
	err = os.Remove(pathname)
	return err
}
