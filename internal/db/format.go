package db

import (
	"log"
	dbModel "yuki-image/internal/db/model"
)

func InsertFormat(format dbModel.Format) (uint64, error) {
	tx := db.Create(&format)
	if tx.Error != nil {
		log.Println(tx.Error)
		return 0, tx.Error
	}
	return SelectFormatIdFromName(format.Name)
}

func SelectFormat(id uint64) (dbModel.Format, error) {
	var format dbModel.Format
	tx := db.First(&format, "id=?", id)
	if tx.Error != nil {
		return dbModel.Format{}, tx.Error
	}
	return format, nil
}

func SelectFormatIdFromName(name string) (uint64, error) {
	var id uint64
	tx := db.Model(dbModel.Format{}).Where("name=?", name).Pluck("id", &id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return id, nil
}

func SelectAllFormat() ([]dbModel.Format, error) {
	var formats []dbModel.Format
	tx := db.Find(&formats)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return formats, nil
}
