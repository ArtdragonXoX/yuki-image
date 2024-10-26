package conf

type ServerConf struct {
	Port  string `yaml:"port"`
	Path  string `yaml:"path"`
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
}
