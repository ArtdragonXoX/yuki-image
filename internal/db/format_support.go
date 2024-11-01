package db

import (
	dbModel "yuki-image/internal/db/model"
)

func InsertFormatSupport(formatSupport dbModel.FormatSupport) error {
	tx := db.Create(&formatSupport)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectFormatSupport(albumId uint64) ([]dbModel.FormatSupport, error) {
	var formats []dbModel.FormatSupport
	tx := db.Model(dbModel.FormatSupport{}).Where("album_id = ?", albumId).Find(&formats)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return formats, nil
}

func DeleteFormatSupport(format dbModel.FormatSupport) error {
	tx := db.Where("album_id=? And format_id =? ", format.AlbumId, format.FormatId).Delete(&dbModel.FormatSupport{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
