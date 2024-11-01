package conf

type DatabaseConf struct {
	MaxConn int  `yaml:"max_open_conns"`
	MaxIdle int  `yaml:"max_idle_conns"`
	Reset   bool `yaml:"reset"`
}
