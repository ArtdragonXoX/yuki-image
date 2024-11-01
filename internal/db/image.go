package db

import (
	"errors"
	"time"
	dbModel "yuki-image/internal/db/model"

	"gorm.io/gorm"
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

func ContainsImageName(name string) (bool, error) {
	var image dbModel.Image
	err := db.Model(dbModel.Image{}).Where("name = ?", name).First(&image).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func ContainsImageKey(key string) (bool, error) {
	var image dbModel.Image
	err := db.Model(dbModel.Image{}).Where("key = ?", key).First(&image).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}
