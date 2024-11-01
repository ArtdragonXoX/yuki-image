package conf

import (
	"fmt"
	"yuki-image/utils"
)

type Config struct {
	DB     DatabaseConf `yaml:"db"`
	Server ServerConf   `yaml:"server"`
	Image  ImageConf    `yaml:"image"`
}

var Conf = &Config{}

func InitConfig() error {
	err := utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	utils.BaseUrl = fmt.Sprintf("%s", Conf.Image.Url)
	utils.KeyLength = Conf.Image.KeyLength
	return nil
}

func WriteConfig() error {
	return utils.WriteYaml(&Conf, "config.yaml")
}
