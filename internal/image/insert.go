package image

import (
	"fmt"
	"regexp"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/internal/image/file"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func Upload(tmpPath string, albumId uint64) (model.Image, error) {
	format := utils.GetImageFormat(tmpPath)
	album, err := db.SelectAlbum(albumId)
	if err != nil {
		return model.Image{}, err
	}
	hash, err := utils.GetImageHash(tmpPath)
	if err != nil {
		return model.Image{}, err
	}
	newFileName := fmt.Sprintf("%s.%s", hash, utils.GetImageFormatName(format))
	newFilePath := fmt.Sprintf("%s/%s", album.Name, newFileName)
	localFilePath := fmt.Sprintf("%s/%s", conf.Conf.Server.Path, newFilePath)
	switch format {
	case model.JPEG:
		file.ManipulateJPEG(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.PNG:
		file.ManipulatePNG(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	case model.GIF:
		file.ManipulateGIF(tmpPath, localFilePath, int(album.MaxHeight), int(album.MaxWidth))
	default:
	}
	size, err := utils.GetImageSize(localFilePath)
	if err != nil {
		return model.Image{}, err
	}

	re := regexp.MustCompile(`[^/]+$`)
	filename := re.FindString(tmpPath)

	image := model.Image{
		Id:         utils.GetRandKey(),
		Name:       newFileName,
		AlbumId:    albumId,
		Pathname:   newFilePath,
		OriginName: filename,
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
