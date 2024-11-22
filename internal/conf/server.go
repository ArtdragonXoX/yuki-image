package conf

import (
	"fmt"
	"strconv"
)

type ServerConf struct {
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (s *ServerConf) Default() {
	s.Port = "7415"
	s.Token, _ = GenerateToken()
}

func (s *ServerConf) Update(new ServerConf) error {
	if err := new.Check(); err != nil {
		return err
	}
	s.Port = new.Port
	s.Token = new.Token
	return nil
}

func (s *ServerConf) Check() error {
	port, err := strconv.Atoi(s.Port) // 尝试将字符串转换为整数
	if err != nil {
		return fmt.Errorf("invalid port number: %s", s.Port) // 如果转换失败，说明不是有效的数字
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("port number out of range: %d", port) // 如果端口号不在合法范围内，返回错误
	}
	return nil
}
