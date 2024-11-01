package db

import (
	"yuki-image/internal/conf"
	dbModel "yuki-image/internal/db/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDataBase() error {
	var err error
	db, err = gorm.Open(sqlite.Open("yuki-image.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&dbModel.Album{}, &dbModel.Format{}, &dbModel.Image{})
	if err != nil {
		return err
	}
	if conf.Conf.DB.Reset {
		InsertFormat(dbModel.Format{Name: "jpeg"})
		InsertFormat(dbModel.Format{Name: "png"})
		InsertFormat(dbModel.Format{Name: "gif"})
		db.Exec("PRAGMA foreign_keys = ON")
	}
	return nil
}
