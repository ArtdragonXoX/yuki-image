package image

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id string) (model.Image, error) {
	return db.SelectImage(id)
}
