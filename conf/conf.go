package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB     DatabaseConf `yaml:"db"`
	Server ServerConf   `yaml:"server"`
}

var Conf = &Config{}

func InitConfig() error {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(Conf)
	if err != nil {
		return err
	}
	return nil
}
