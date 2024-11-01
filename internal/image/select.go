package image

import (
	"regexp"
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(key string) (model.Image, error) {
	dbimage, err := db.SelectImage(key)
	if err != nil {
		return model.Image{}, err
	}
	var image model.Image
	image.FromDBModel(dbimage)
	return image, nil
}

func SelectFromUrl(url string) (model.Image, error) {
	pattern := "[\u4e00-\u9fa5a-zA-Z0-9]+/[a-zA-Z0-9]+\\.[a-zA-Z0-9]+$"
	match := regexp.MustCompile(pattern).FindString(url)
	key, err := db.SelectImageKeyFromPath(match)
	if err != nil {
		return model.Image{}, err
	}
	return Select(key)
}
