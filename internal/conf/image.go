package conf

import "errors"

type ImageConf struct {
	KeyLength            int    `yaml:"key_length" json:"key_length"`
	MaxSize              int    `yaml:"max_size" json:"max_size"`
	Path                 string `yaml:"path" json:"path"`
	AutoDeleteTmp        bool   `yaml:"auto_delete_tmp" json:"auto_delete_tmp"`
	ImageListDefalutSize int    `yaml:"image_list_defalut_size" json:"image_list_defalut_size"`
	CompressionQuality   int    `yaml:"compression_quality" json:"compression_quality"`
	Url                  string `yaml:"url" json:"url"`
}

func (i *ImageConf) Default() {
	i.KeyLength = 10
	i.MaxSize = 20
	i.Path = "./localimage"
	i.AutoDeleteTmp = true
	i.ImageListDefalutSize = 10
	i.CompressionQuality = 6
	i.Url = "http://127.0.0.1:7415"
}

func (i *ImageConf) Update(new ImageConf) error {
	if err := new.Check(); err != nil {
		return err
	}
	i.KeyLength = new.KeyLength
	i.MaxSize = new.MaxSize
	i.Path = new.Path
	i.AutoDeleteTmp = new.AutoDeleteTmp
	i.ImageListDefalutSize = new.ImageListDefalutSize
	i.CompressionQuality = new.CompressionQuality
	return nil
}

func (i *ImageConf) Check() error {
	if i.KeyLength <= 0 {
		return errors.New("key_length must be greater than 0")
	}
	if i.MaxSize <= 0 {
		return errors.New("max_size must be greater than 0")
	}
	if i.ImageListDefalutSize <= 0 {
		return errors.New("image_list_defalut_size must be greater than 0")
	}
	if i.CompressionQuality < 1 || i.CompressionQuality > 6 {
		return errors.New("compression_quality must be between 1 and 6")
	}
	return nil
}
