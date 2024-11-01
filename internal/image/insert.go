package image

import (
	"errors"
	"fmt"
	"log"
	"os"
	ialbum "yuki-image/internal/album"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	imageio "yuki-image/internal/image/file"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func Upload(tmpPath string, oname string, albumId uint64) (model.Image, error) {
	file, err := os.Open(tmpPath)
	if err != nil {
		return model.Image{}, err
	}
	defer file.Close()
	buff := make([]byte, 512)

	_, err = file.Read(buff)
	if err != nil {
		return model.Image{}, err
	}

	format := utils.GetImageFormat(buff)
	album, err := ialbum.Select(albumId)
	if err != nil {
		return model.Image{}, err
	}
	albumFormat, err := ialbum.SelectFormatSupport(albumId)
	if err != nil {
		return model.Image{}, err
	}
	var formats []string
	for _, v := range albumFormat {
		formats = append(formats, v.Name)
	}
	if !utils.Contains(formats, utils.GetImageFormatName(format)) {
		return model.Image{}, errors.New("format not supported")
	}
	var hash string
	for {
		hash, err = utils.GetByteHash(buff)
		if err != nil {
			return model.Image{}, err
		}
		v, err := db.ContainsImageName(hash)
		if err != nil {
			return model.Image{}, err
		}
		if !v {
			break
		}
	}
	newFileName := fmt.Sprintf("%s.%s", hash, utils.GetImageFormatName(format))
	newFilePath := fmt.Sprintf("%s/%s", album.Name, newFileName)
	localFilePath := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, newFilePath)

	switch format {
	case model.JPEG:
		err = imageio.ManipulateJPEG(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.PNG:
		err = imageio.ManipulatePNG(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.GIF:
		err = imageio.ManipulateGIF(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	default:
	}
	if err != nil {
		log.Println(err)
		return model.Image{}, err
	}

	size, err := utils.GetImageSize(localFilePath)
	if err != nil {
		return model.Image{}, err
	}

	var key string
	for {
		key = utils.GetRandKey()
		v, err := db.ContainsImageKey(key)
		if err != nil {
			return model.Image{}, err
		}
		if !v {
			break
		}
	}

	image := model.Image{
		Key:        key,
		Name:       newFileName,
		AlbumId:    albumId,
		Pathname:   newFilePath,
		OriginName: oname,
		Size:       size,
		Mimetype:   utils.GetImageFormatName(format),
	}
	err = db.InsertImage(image.ToDBModel())
	if err != nil {
		return model.Image{}, err
	}
	image, err = Select(image.Key)
	if err != nil {
		return model.Image{}, err
	}
	image.Url = utils.GetImageUrl(image)
	return image, nil
}
