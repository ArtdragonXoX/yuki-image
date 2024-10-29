package conf

type ImageConf struct {
	MaxSize              int    `yaml:"max_size"`
	Path                 string `yaml:"path"`
	ImageListDefalutSize int    `yaml:"image_list_defalut_size"`
	CompressionQuality   int    `yaml:"compression_quality"`
	Url                  string `yaml:"url"`
}
