package db

import (
	"log"
	"time"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func InsertImage(image model.Image) error {
	sql := "INSERT INTO tbl_image (id, name,  album_id, pathname, origin_name, size, mimetype, time) VALUES (?, ?,  ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(image.Id, image.Name, image.AlbumId, image.Pathname, image.OriginName, image.Size, image.Mimetype, time)
	if imagetmp, err := utils.PrettyStruct(image); err != nil {
		log.Println("Pretty struct err:", err)
		log.Println(sql, image)
	} else {
		log.Println(imagetmp)
	}
	if err != nil {
		return err
	}

	return nil
}

func SelectImage(id string) (model.Image, error) {
	var image model.Image
	sql := "SELECT * FROM tbl_image WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return image, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&image.Id, &image.Name, &image.AlbumId, &image.Pathname, &image.OriginName, &image.Size, &image.Mimetype, &image.Time)
	image.Url = utils.GetImageUrl(image)
	return image, err
}

func DeleteImage(id string) error {
	sql := "DELETE FROM tbl_image WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
