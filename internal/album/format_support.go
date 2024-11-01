package album

import (
	"yuki-image/internal/db"
	iformat "yuki-image/internal/format"
	"yuki-image/internal/model"
)

func InsertFormatSupport(formatSupport model.FormatSupport) error {
	formatSupport, err := GetFormatSupportId(formatSupport)
	if err != nil {
		return err
	}
	return db.InsertFormatSupport(formatSupport.ToDBModel())
}
func DeleteFormatSupport(formatSupport model.FormatSupport) error {
	formatSupport, err := GetFormatSupportId(formatSupport)
	if err != nil {
		return err
	}
	return db.DeleteFormatSupport(formatSupport.ToDBModel())
}

func GetFormatSupportId(formatSupport model.FormatSupport) (model.FormatSupport, error) {
	var err error
	if formatSupport.AlbumId <= 0 {
		formatSupport.AlbumId, err = SelectIdFromName(formatSupport.AlbumName)
		if err != nil {
			return model.FormatSupport{}, err
		}
	}
	if formatSupport.FormatId <= 0 {
		formatSupport.FormatId, err = iformat.SelectIdFromName(formatSupport.FormatName)
		if err != nil {
			return model.FormatSupport{}, err
		}
	}
	return formatSupport, nil
}

func SelectFormatSupportFromName(name string) ([]model.Format, error) {
	id, err := SelectIdFromName(name)
	if err != nil {
		return nil, err
	}
	return SelectFormatSupport(id)

}

func SelectFormatSupport(albumId uint64) ([]model.Format, error) {
	dbformatsupport, err := db.SelectFormatSupport(albumId)
	if err != nil {
		return nil, err
	}
	var formats []model.Format
	for _, dbformat := range dbformatsupport {
		format, err := iformat.Select(dbformat.FormatId)
		if err != nil {
			return nil, err
		}
		formats = append(formats, format)
	}
	return formats, nil
}
