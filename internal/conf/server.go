package conf

type ServerConf struct {
	Port  string `yaml:"port"`
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
}
