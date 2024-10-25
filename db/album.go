package db

import (
	"log"
	"time"
	"yuki-image/model"
	"yuki-image/utils"
)

func InsertAlbum(album model.Album) (uint64, error) {
	sql := "INSERT INTO tbl_album (name,max_height,max_width,update_time,create_time) VALUES (?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	time := time.Now().Format("2006-01-02 15:04:05")
	result, err := stmt.Exec(album.Name, album.MaxHeight, album.MaxWidth, time, time)
	if albumtmp, err := utils.PrettyStruct(album); err != nil {
		log.Println("Pretty struct err:", err)
		log.Println(sql, album)
	} else {
		log.Println(albumtmp)
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

func SelectAlbum(id uint64) (model.Album, error) {
	var album model.Album
	sql := "SELECT * FROM tbl_album WHERE id = ?"
	log.Println(sql, id)
	err := db.QueryRow(sql, id).Scan(&album.Id, &album.Name, &album.MaxHeight, &album.MaxWidth, &album.UpdateTime, &album.CreateTime)
	if err != nil {
		return model.Album{}, err
	}
	format_support, err := SelectFormatSupport(album.Id)
	if err != nil {
		return model.Album{}, err
	}
	album.FormatSupport = format_support
	return album, err
}

func SelectAllAlbum() ([]model.Album, error) {
	sql := "SELECT id FROM tbl_album"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albums := make([]model.Album, 0)
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		album, err := SelectAlbum(id)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}
