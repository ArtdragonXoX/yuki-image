package format

import (
	"yuki-image/internal/db"
	"yuki-image/internal/model"
)

func Select(id uint64) (model.Format, error) {
	dbformat, err := db.SelectFormat(id)
	if err != nil {
		return model.Format{}, err
	}
	var format model.Format
	format.FromDBModel(dbformat)
	return format, nil
}

func SelectIdFromName(name string) (uint64, error) {
	return db.SelectFormatIdFromName(name)
}

func SelectFormatFromName(name string) (model.Format, error) {
	id, err := SelectIdFromName(name)
	if err != nil {
		return model.Format{}, err
	}
	return Select(id)
}

func SelectAll() ([]model.Format, error) {
	dbformats, err := db.SelectAllFormat()
	if err != nil {
		return []model.Format{}, err
	}
	var formats []model.Format
	for _, dbformat := range dbformats {
		var format model.Format
		format.FromDBModel(dbformat)
		formats = append(formats, format)
	}
	return formats, nil
}
