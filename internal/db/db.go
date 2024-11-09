package db

import (
	"log"
	"os"
	"yuki-image/internal/conf"
	dbModel "yuki-image/internal/db/model"
	"yuki-image/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var dbFile = "./database/yuki-image.db"

func InitDataBase() error {
	if conf.Conf.DB.Reset {
		err := ResetDB()
		if err != nil {
			return err
		}
		conf.Conf.DB.Reset = false
		utils.WriteYaml(conf.Conf, "config.yaml")
	}
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(conf.Conf.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.Conf.DB.MaxConn)
	db.Exec("PRAGMA foreign_keys = ON")
	log.Println("database init success!")
	return nil
}

func ResetDB() error {
	utils.EnsureDir(dbFile)
	v, err := utils.IsFileExists(dbFile)
	if err != nil {
		return err
	}
	if v {
		os.Remove(dbFile)
	}

	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&dbModel.Album{}, &dbModel.Format{}, &dbModel.Image{})
	if err != nil {
		return err
	}
	InsertFormat(dbModel.Format{Name: "jpeg"})
	InsertFormat(dbModel.Format{Name: "png"})
	InsertFormat(dbModel.Format{Name: "gif"})
	log.Println("database reset success!")
	return nil
}
