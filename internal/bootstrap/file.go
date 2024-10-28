package bootstrap

import (
	"yuki-image/internal/conf"
	"yuki-image/utils"
)

func InitFile() error {
	err := utils.EnsureDir("tmp")
	if err != nil {
		return err
	}
	err = utils.EnsureDir(conf.Conf.Image.Path)
	if err != nil {
		return err
	}
	return nil
}
