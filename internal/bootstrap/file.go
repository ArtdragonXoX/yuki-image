package bootstrap

import (
	"yuki-image/internal/conf"
	"yuki-image/internal/tmp"
	"yuki-image/utils"
)

func InitFile() error {
	err := utils.EnsureDir(conf.Conf.Image.Path)
	if err != nil {
		return err
	}
	tmp.TmpPath = conf.Conf.Image.Path + "/tmp"
	err = utils.EnsureDir(tmp.TmpPath)
	if err != nil {
		return err
	}
	return nil
}
