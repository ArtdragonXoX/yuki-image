package db

import (
	"time"
	dbModel "yuki-image/internal/db/model"
)

func InsertImage(image dbModel.Image) error {
	time := time.Now()
	image.CreateTime = time
	tx := db.Create(&image)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectImage(key string) (dbModel.Image, error) {
	var image dbModel.Image
	tx := db.First(&image, "key = ?", key)
	if tx.Error != nil {
		return dbModel.Image{}, tx.Error
	}
	return image, nil
}

func SelectImageKeyFromPath(pathname string) (string, error) {
	var key string
	tx := db.Model(dbModel.Image{}).Where("pathname = ?", pathname).Pluck("key", &key)
	if tx.Error != nil {
		return "", tx.Error
	}
	return key, nil
}
func DeleteImage(key string) error {
	tx := db.Delete(&dbModel.Image{}, "key = ?", key)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
