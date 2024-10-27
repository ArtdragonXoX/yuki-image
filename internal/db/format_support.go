package db

import (
	"log"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func InsertFormatSupport(format model.FormatSupport) error {
	sql := "INSERT INTO tbl_format_support (format_id,album_id) VALUES (?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(format.FormatId, format.AlbumId)
	if formattmp, err := utils.PrettyStruct(format); err != nil {
		log.Println("Pretty struct err:", err)
		log.Println(sql, format)
	} else {
		log.Println(formattmp)
	}
	if err != nil {
		return err
	}
	return nil
}

func SelectFormatSupport(albumId uint64) ([]model.Format, error) {
	var formats []model.Format
	sql := "SELECT format_id FROM tbl_format_support WHERE album_id=?"
	rows, err := db.Query(sql, albumId)
	if err != nil {
		return formats, err
	}
	for rows.Next() {
		var format_id uint64
		err = rows.Scan(&format_id)
		format, err := SelectFormat(format_id)
		if err != nil {
			return nil, err
		}
		formats = append(formats, format)
	}
	defer rows.Close()
	return formats, nil
}

func DeleteFormatSupport(format model.FormatSupport) error {
	sql := "DELETE FROM tbl_format_support WHERE format_id=? AND album_id=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(format.FormatId, format.AlbumId)
	if err != nil {
		return err
	}
	return nil
}
