package image

import (
	"regexp"
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id string) (model.Image, error) {
	return db.SelectImage(id)
}

func SelectFromUrl(url string) (model.Image, error) {
	pattern := "[\u4e00-\u9fa5a-zA-Z0-9]+/[a-zA-Z0-9]+\\.[a-zA-Z0-9]+$"
	match := regexp.MustCompile(pattern).FindString(url)
	id, err := db.SelectImageIdFromPath(match)
	if err != nil {
		return model.Image{}, err
	}
	return db.SelectImage(id)
}
