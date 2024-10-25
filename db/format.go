package db

import (
	"log"
	"yuki-image/model"
	"yuki-image/utils"
)

func InsertFormat(format model.Format) (uint64, error) {
	sql := "INSERT INTO tbl_format(id,name) VALUES (?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(format.Id, format.Name)
	if formattmp, err := utils.PrettyStruct(format); err != nil {
		log.Println("Pretty struct err:", err)
		log.Println(sql, format)
	} else {
		log.Println(formattmp)
	}
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func SelectFormat(id uint64) (model.Format, error) {
	sql := "SELECT * FROM tbl_format WHERE id=?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return model.Format{}, err
	}
	var format model.Format
	err = stmt.QueryRow(id).Scan(&format.Id, &format.Name)
	if err != nil {
		return model.Format{}, err
	}
	return format, nil
}

func SelectAllFormat() ([]model.Format, error) {
	sql := "SELECT * FROM tbl_format"
	rows, err := db.Query(sql)
	if err != nil {
		return []model.Format{}, err
	}
	formats := make([]model.Format, 0)
	for rows.Next() {
		var f model.Format
		err := rows.Scan(&f.Id, &f.Name)
		if err != nil {
			return []model.Format{}, err
		}
		formats = append(formats, f)
	}
	return formats, nil
}
