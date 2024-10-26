package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yuki-image/conf"
	"yuki-image/model"
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
	err := CheckAlbum()
	if err != nil {
		return err
	}
	err = CheckImage()
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
	exist := 0
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s is exist.\n", album, column))
		}
		exist++
	}

	if exist != len(columns) {
		return errors.New(fmt.Sprintf("Table %s is not complete.\n", album))
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
		return errors.New(fmt.Sprintf("Table %s is exist.\n", image))
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
	exist := 0
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s is exist.\n", image, column))
		}
		exist++
	}
	if exist != len(columns) {
		return errors.New(fmt.Sprintf("Table %s is not complete.\n", image))
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
	exist := 0
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s is exist.\n", format, column))
		}
		exist++
	}
	if exist != len(columns) {
		return errors.New(fmt.Sprintf("Table %s is not complete.\n", format))
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
	exist := 0
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return err
		}
		if !utils.Contains(columns, column) {
			return errors.New(fmt.Sprintf("Table %s column %s is exist.\n", format_support, column))
		}
		exist++
	}
	if exist != len(columns) {
		return errors.New(fmt.Sprintf("Table %s is not complete.\n", format_support))
	}
	return nil
}

func ResetTable() error {
	err := DropAllTable()
	if err != nil {
		return err
	}
	err = CreateAllTable()
	if err != nil {
		return err
	}
	return nil
}

func ResetAlbum() error {
	err := DropAlbum()
	if err != nil {
		return err
	}
	err = CreateAlbum()
	if err != nil {
		return err
	}
	return nil
}

func ResetImage() error {
	err := DropImage()
	if err != nil {
		return err
	}
	err = CreateImage()
	if err != nil {
		return err
	}
	return nil
}

func ResetFormat() error {
	err := DropFormat()
	if err != nil {
		return err
	}
	err = CreateFormat()
	if err != nil {
		return err
	}
	return nil
}

func ResetFormatSupport() error {
	err := DropFormatSupport()
	if err != nil {
		return err
	}
	err = CreateFormatSupport()
	if err != nil {
		return err
	}
	return nil
}

func DropAllTable() error {
	err := DropFormatSupport()
	if err != nil {
		return err
	}
	err = DropImage()
	if err != nil {
		return err
	}
	err = DropAlbum()
	if err != nil {
		return err
	}
	err = DropFormat()
	if err != nil {
		return err
	}
	return nil
}

func DropAlbum() error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", conf.Conf.DB.Name, "tbl_album")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func DropImage() error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", conf.Conf.DB.Name, "tbl_image")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func DropFormat() error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", conf.Conf.DB.Name, "tbl_format")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func DropFormatSupport() error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", conf.Conf.DB.Name, "tbl_format_support")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateAllTable() error {
	err := CreateAlbum()
	if err != nil {
		return err
	}
	err = CreateImage()
	if err != nil {
		return err
	}
	err = CreateFormat()
	if err != nil {
		return err
	}
	err = CreateFormatSupport()
	if err != nil {
		return err
	}
	return nil
}

func CreateAlbum() error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(255) UNIQUE NOT NULL,max_height INT NOT NULL, max_width INT NOT NULL, update_time TIMESTAMP NOT NULL, create_time TIMESTAMP NOT NULL)", conf.Conf.DB.Name, "tbl_album")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateImage() error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (id VARCHAR(255) NOT NULL PRIMARY KEY, name VARCHAR(255) UNIQUE NOT NULL, album_id INT UNSIGNED NOT NULL, pathname VARCHAR(255) NOT NULL, origin_name VARCHAR(255) NOT NULL, size INT UNSIGNED NOT NULL,  mimetype VARCHAR(255) NOT NULL, time timestamp NOT NULL, FOREIGN KEY (album_id) REFERENCES %s(%s) ON DELETE CASCADE ON UPDATE CASCADE)", conf.Conf.DB.Name, "tbl_image", "tbl_album", "id")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func CreateFormat() error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(255) UNIQUE NOT NULL)", conf.Conf.DB.Name, "tbl_format")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	InsertFormat(model.Format{Id: 1, Name: "jpeg"})
	InsertFormat(model.Format{Id: 2, Name: "png"})
	InsertFormat(model.Format{Id: 3, Name: "gif"})
	return nil
}

func CreateFormatSupport() error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (format_id INT UNSIGNED NOT NULL, album_id INT UNSIGNED NOT NULL,PRIMARY KEY (format_id, album_id), FOREIGN KEY (format_id) REFERENCES %s(%s) ON DELETE CASCADE ON UPDATE CASCADE ,FOREIGN KEY (album_id) REFERENCES %s(%s) ON DELETE CASCADE ON UPDATE CASCADE)", conf.Conf.DB.Name, "tbl_format_support", "tbl_format", "id", "tbl_album", "id")
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
