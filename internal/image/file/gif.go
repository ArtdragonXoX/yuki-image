package file

import (
	"log"
	"yuki-image/internal/conf"
	"yuki-image/utils"

	"github.com/disintegration/imaging"
)

func ManipulateGIF(tmpPath string, path string, max_height int, max_width int) error {
	img, err := imaging.Open(tmpPath)
	if err != nil {
		return err
	}
	bounds := img.Bounds()

	width := bounds.Dx()
	height := bounds.Dy()
	if width > max_width || height > max_height {
		if height > max_height {
			width = int(float64(max_height) * float64(width) / float64(height))
			height = max_height
		}
		if width > max_width {
			height = int(float64(max_width) * float64(height) / float64(width))
			width = max_width
		}
		img = imaging.Resize(img, width, height, utils.GetResampleFilter(conf.Conf.Image.CompressionQuality))
		log.Println("resize", width, height)
	}

	err = imaging.Save(img, path)
	if err != nil {
		return err
	}

	return nil
}
