package album

import (
	"yuki-image/internal/db"
	iformat "yuki-image/internal/format"
	"yuki-image/internal/model"
)

func InsertFormatSupport(formatSupport model.FormatSupport) error {
	albumId := formatSupport.AlbumId
	formatId := formatSupport.FormatId
	var err error
	if albumId == 0 {
		albumId, err = SelectIdFromName(formatSupport.AlbumName)
		if err != nil {
			return err
		}
	}
	if formatId == 0 {
		formatId, err = iformat.SelectIdFromName(formatSupport.FormatName)
		if err != nil {
			return err
		}
	}
	return db.InsertFormatSupport(model.FormatSupport{
		AlbumId:  albumId,
		FormatId: formatId,
	})
}
func DeleteFormatSupport(formatSupport model.FormatSupport) error {
	albumId := formatSupport.AlbumId
	formatId := formatSupport.FormatId
	var err error
	if albumId == 0 {
		albumId, err = SelectIdFromName(formatSupport.AlbumName)
		if err != nil {
			return err
		}
	}
	if formatId == 0 {
		formatId, err = iformat.SelectIdFromName(formatSupport.FormatName)
		if err != nil {
			return err
		}
	}
	return db.DeleteFormatSupport(model.FormatSupport{
		AlbumId:  albumId,
		FormatId: formatId,
	})
}

func SelectFormatSupportFromName(name string) ([]model.Format, error) {
	id, err := db.SelectAlbumIdFromName(name)
	if err != nil {
		return nil, err
	}
	return db.SelectFormatSupport(id)

}

func SelectFormatSupport(albumId uint64) ([]model.Format, error) {
	return db.SelectFormatSupport(albumId)
}
