package file

import (
	"image"
	"image/png"
	"os"
	"yuki-image/utils"
)

func ManipulatePNG(tmpPath string, path string, max_height int, max_width int) error {
	file, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	bounds := img.Bounds()

	width := bounds.Dx()
	height := bounds.Dy()
	if width > max_width || height > max_height {
		if height > max_height {
			height = max_height
			width = int(float64(max_height) * float64(width) / float64(height))
		} else {
			width = max_width
			height = int(float64(max_width) * float64(height) / float64(width))
		}
		img = utils.ResizeImage(img, width, height)
	}

	outFile, err := os.Create(path)

	if err != nil {
		return err
	}
	defer outFile.Close()
	err = png.Encode(outFile, img)
	if err != nil {
		return err
	}
	return nil
}
