package db

import (
	"log"
	"time"
	"yuki-image/internal/model"
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

func UpdateAlbum(album model.Album) error {
	sql := "UPDATE tbl_album SET name = ?,max_height = ?,max_width = ?,update_time=? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	album_, err := SelectAlbum(album.Id)
	if err != nil {
		return err
	}
	if album.Name == "" {
		album.Name = album_.Name
	}
	if album.MaxHeight == 0 {
		album.MaxHeight = album_.MaxHeight
	}
	if album.MaxWidth == 0 {
		album.MaxWidth = album_.MaxWidth
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(album.Name, album.MaxHeight, album.MaxWidth, time, album.Id)
	if err != nil {
		return err
	}
	return nil
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
	total, err := GetAlbumImageTotal(id)
	if err != nil {
		return album, err
	}
	album.Image = model.ImageList{Total: total}

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

func SelectImageFromAlbum(albumId uint64, page uint64, size uint64) (model.ImageList, error) {
	var images []model.Image
	var imageList model.ImageList

	sql := "SELECT id FROM tbl_image WHERE album_id = ? LIMIT ?,?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return model.ImageList{}, err
	}
	defer stmt.Close()
	rowNum := (page - 1) * size
	rows, err := stmt.Query(albumId, rowNum, size)
	for rows.Next() {
		var imageId string
		err = rows.Scan(&imageId)
		if err != nil {
			return model.ImageList{}, err
		}
		image, err := SelectImage(imageId)
		if err != nil {
			return model.ImageList{}, err
		}
		images = append(images, image)
	}
	imageList.Image = images
	imageList.Total, err = GetAlbumImageTotal(albumId)
	if err != nil {
		return model.ImageList{}, err
	}
	imageList.Page = page
	imageList.Size = size
	return imageList, err
}

func GetAlbumImageTotal(albumId uint64) (uint64, error) {
	var count uint64
	sql := "SELECT COUNT(*) FROM tbl_image WHERE album_id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(albumId).Scan(&count)
	return count, err
}
