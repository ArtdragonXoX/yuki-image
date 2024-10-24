package db

import (
	"database/sql"
	"log"
	"yuki-image/conf"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDataBase() error {
	data := conf.Conf.DB
	var err error

	dsn := data.User + ":" + data.Pwd + "@tcp(" + data.Host + ":" + data.Port + ")/" + data.Name
	db, err = sql.Open("mysql", dsn)
	log.Println("Connecting to MySQL:", dsn)

	if err != nil {
		log.Println("Open database error!", err)
		return err
	}
	db.SetMaxIdleConns(data.MaxIdle)
	db.SetMaxOpenConns(data.MaxConn)
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database!")
		return err
	}

	log.Println("Successfully connected to MySQL!")
	return nil
}
