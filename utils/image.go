package utils

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	imgext "github.com/shamsher31/goimgext"
	"golang.org/x/image/draw"
)

const (
	JPEG uint64 = 1
	PNG  uint64 = 2
	GIF  uint64 = 3
)

var BaseUrl string

func GetResampleFilter(i int) imaging.ResampleFilter {
	switch i {
	case 1:
		return imaging.NearestNeighbor
	case 2:
		return imaging.Box
	case 3:
		return imaging.Linear
	case 4:
		return imaging.MitchellNetravali
	case 5:
		return imaging.CatmullRom
	case 6:
		return imaging.Lanczos
	default:
		return imaging.Box
	}

}

func GetImageSize(filepath string) (uint64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err // 如果发生错误，返回0和错误
	}
	// 返回文件大小
	return uint64(fileInfo.Size()), nil
}

func GetImageFormatName(format uint64) string {
	switch format {
	case JPEG:
		return "jpeg"
	case PNG:
		return "png"
	case GIF:
		return "gif"

	default:
		return ""
	}
}
func GetImageFormat(buff []byte) uint64 {
	filetype := http.DetectContentType(buff)
	ext := imgext.Get()
	var datatype string
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			datatype = filetype
			break
		}
	}
	log.Println(datatype)
	switch datatype {
	case `image/jpeg`:
		return JPEG
	case `image/png`:
		return PNG
	case `image/gif`:
		return GIF
	default:
		return 0
	}
}

func ResizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

func GetImageUrl(pathname string) string {
	return fmt.Sprintf("%s/i/%s", BaseUrl, pathname)
}
