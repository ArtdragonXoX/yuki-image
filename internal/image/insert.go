package image

import (
	"fmt"
	"log"
	"os"
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
	buff := make([]byte, 512)

	_, err = file.Read(buff)
	if err != nil {
		return model.Image{}, err
	}

	format := utils.GetImageFormat(buff)
	album, err := db.SelectAlbum(albumId)
	if err != nil {
		return model.Image{}, err
	}
	hash, err := utils.GetByteHash(buff)
	if err != nil {
		return model.Image{}, err
	}
	newFileName := fmt.Sprintf("%s.%s", hash, utils.GetImageFormatName(format))
	newFilePath := fmt.Sprintf("%s/%s", album.Name, newFileName)
	localFilePath := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, newFilePath)

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
	err = os.Remove(tmpPath)
	if err != nil {
		return model.Image{}, err
	}

	size, err := utils.GetImageSize(localFilePath)
	if err != nil {
		return model.Image{}, err
	}

	image := model.Image{
		Id:         utils.GetRandKey(),
		Name:       newFileName,
		AlbumId:    albumId,
		Pathname:   newFilePath,
		OriginName: oname,
		Size:       size,
		Mimetype:   utils.GetImageFormatName(format),
	}
	err = db.InsertImage(image)
	if err != nil {
		return model.Image{}, err
	}
	image, err = db.SelectImage(image.Id)
	if err != nil {
		return model.Image{}, err
	}
	image.Url = utils.GetImageUrl(image)
	return image, nil
}
