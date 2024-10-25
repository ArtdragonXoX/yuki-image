package conf

type ServerConf struct {
	Port  string `yaml:"port"`
	Path  string `yaml:"path"`
	Token string `yaml:"token"`
}
