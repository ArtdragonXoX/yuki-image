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
	err = db.AutoMigrate(&dbModel.Album{}, &dbModel.Format{}, &dbModel.Image{})
	if err != nil {
		return err
	}

	// 确保基本格式(jpeg, png, gif)存在
	// 如果表为空，它们将按指定顺序创建
	// 这会影响它们的自增ID
	defaultFormats := []dbModel.Format{
		{Name: "jpeg"},
		{Name: "png"},
		{Name: "gif"},
	}

	for _, format := range defaultFormats {
		// FirstOrCreate查找符合条件(format.Name)的第一条记录
		// 如果未找到则用给定条件创建新记录
		// 第二个参数(format)用于在记录不存在时创建
		if tx := db.Where(dbModel.Format{Name: format.Name}).FirstOrCreate(&dbModel.Format{Name: format.Name}); tx.Error != nil {
			// 记录错误但不一定panic
			// 因为核心数据库初始化可能已成功
			// 根据需求，这里可以返回错误
			log.Printf("Error ensuring format %s: %v\n", format.Name, tx.Error)
		}
	}

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
