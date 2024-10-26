package bootstrap

import (
	"yuki-image/conf"
	"yuki-image/utils"
)

func InitFile() error {
	err := utils.EnsureDir("tmp")
	if err != nil {
		return err
	}
	err = utils.EnsureDir(conf.Conf.Server.Path)
	if err != nil {
		return err
	}
	return nil
}
