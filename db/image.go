package db

import (
	"log"
	"time"
	"yuki-image/model"
	"yuki-image/utils"
)

func InsertImage(image model.Image) error {
	sql := "INSERT INTO tbl_image (id, name, url, album, pathname, origin_name, size, mimetype, time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(image.Id, image.Name, image.Url, image.Album, image.Pathname, image.OriginName, image.Size, image.Mimetype, time)
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
