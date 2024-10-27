package conf

type ServerConf struct {
	Port                 string `yaml:"port"`
	Path                 string `yaml:"path"`
	Host                 string `yaml:"host"`
	ImageListDefalutSize int    `yaml:"image_list_defalut_size"`
	Token                string `yaml:"token"`
}
