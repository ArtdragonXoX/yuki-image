package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yuki-image/conf"
	"yuki-image/utils"

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

	err = CheckTable()
	if err != nil {
		log.Println("Error checking table!", err)
		return err
	}

	return nil
}

func CheckTable() error {
	err := CheckImage()
	if err != nil {
		return err
	}
	err = CheckAlbum()
	if err != nil {
		return err
	}
	err = CheckFormat()
	if err != nil {
		return err
	}
	err = CheckFormatSupport()
	if err != nil {
		return err
	}
	return nil
}

func CheckImage() error {
	image := "tbl_image"
	sql := fmt.Sprintf("SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, image)
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return errors.New(fmt.Sprintf("Table %s does not exist.\n", image))
	}
	id := "id"
	name := "name"
	album_id := "album_id"
	pathname := "pathname"
	origin_name := "origin_name"
	size := "size"
	mimetype := "mimetype"
	time := "time"
	var columns []string = []string{id, name, album_id, pathname, origin_name, size, mimetype, time}
	sql = fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, image)
	rows, err = db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s does not exist.\n", image, column))
		}
	}

	return nil
}

func CheckAlbum() error {
	album := "tbl_album"
	sql := fmt.Sprintf("SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, album)
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	if !rows.Next() {
		return errors.New(fmt.Sprintf("Table %s does not exist.\n", album))
	}

	id := "id"
	name := "name"
	max_height := "max_height"
	max_width := "max_width"
	update_time := "update_time"
	create_time := "create_time"
	var columns []string = []string{id, name, max_height, max_width, update_time, create_time}

	sql = fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, album)
	rows, err = db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s does not exist.\n", album, column))
		}
	}

	return nil
}

func CheckFormat() error {
	format := "tbl_format"
	sql := fmt.Sprintf("SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, format)
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	id := "id"
	name := "name"
	var columns []string = []string{id, name}

	sql = fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, format)
	rows, err = db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s does not exist.\n", format, column))
		}
	}

	return nil
}

func CheckFormatSupport() error {
	format_support := "tbl_format_support"
	sql := fmt.Sprintf("SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, format_support)
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	format_id := "format_id"
	album_id := "album_id"
	var columns []string = []string{format_id, album_id}

	sql = fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s'", conf.Conf.DB.Name, format_support)
	rows, err = db.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s does not exist.\n", format_support, column))
		}
	}

	return nil
}
