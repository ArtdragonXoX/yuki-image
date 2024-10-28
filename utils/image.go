package utils

import (
	"crypto/md5"
	"fmt"
	"image"
	"log"
	"os"
	"time"
	"yuki-image/internal/model"

	imgtype "github.com/shamsher31/goimgtype"
	"golang.org/x/exp/rand"
	"golang.org/x/image/draw"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var BaseUrl string

func GetImageFormatName(format uint64) string {
	switch format {
	case model.JPEG:
		return "jpeg"
	case model.PNG:
		return "png"
	case model.GIF:
		return "gif"

	default:
		return ""
	}
}

func GetImageFormat(file_name string) uint64 {
	datatype, err := imgtype.Get(file_name)
	if err != nil {
		return 0
	}
	log.Println(datatype)
	switch datatype {
	case `image/jpeg`:
		return model.JPEG
	case `image/png`:
		return model.PNG
	case `image/gif`:
		return model.GIF
	default:
		return 0
	}
}

func GetImageHash(file_name string) (string, error) {
	imageData, err := os.ReadFile(file_name)
	if err != nil {
		return "", err
	}

	timestamp := time.Now().UnixNano()
	timestampBytes := []byte(fmt.Sprintf("%d", timestamp))

	dataTOHash := append(imageData, timestampBytes...)
	hash := md5.Sum(dataTOHash)
	hashHex := fmt.Sprintf("%x", hash)
	return hashHex, nil
}

func GetImageUrl(image model.Image) string {
	return fmt.Sprintf("%s/i/%s", BaseUrl, image.Pathname)
}

func GetRandKey() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]rune, 6)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

func GetImageSize(filepath string) (uint64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err // 如果发生错误，返回0和错误
	}

	// 返回文件大小
	return uint64(fileInfo.Size()), nil
}

func ResizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}
