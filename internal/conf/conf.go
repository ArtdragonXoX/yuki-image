package conf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"yuki-image/utils"
)

type Config struct {
	Server ServerConf   `yaml:"server" json:"server"`
	DB     DatabaseConf `yaml:"db" json:"db"`
	Image  ImageConf    `yaml:"image" json:"image"`
}

var Conf = &Config{}

func InitConfig() error {
	err := utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	utils.BaseUrl = &Conf.Image.Url
	utils.KeyLength = &Conf.Image.KeyLength
	utils.Secret = &Conf.Server.Token
	return nil
}

func WriteConfig() error {
	return utils.WriteYaml(&Conf, "config.yaml")
}

func (c *Config) Update(new Config) error {
	if err := new.Check(); err != nil {
		return err
	}
	if err := c.DB.Update(new.DB); err != nil {
		return err
	}
	if err := c.Server.Update(new.Server); err != nil {
		return err
	}
	if err := c.Image.Update(new.Image); err != nil {
		return err
	}
	return nil
}

func (c *Config) Check() error {
	if err := c.DB.Check(); err != nil {
		return err
	}
	if err := c.Server.Check(); err != nil {
		return err
	}
	if err := c.Image.Check(); err != nil {
		return err
	}
	return nil
}

func GetToken() string {
	return Conf.Server.Token
}

func UpdateToken(token string) {
	Conf.Server.Token = token
	WriteConfig()
}

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(append([]byte(fmt.Sprintf("%d", time.Now().UnixNano())), bytes...))
	token := hex.EncodeToString(hash[:])
	tokenWithUppercase := utils.MakeTokenWithUppercase(token)
	return tokenWithUppercase, nil
}
