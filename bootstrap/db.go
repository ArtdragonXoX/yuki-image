package bootstrap

import "yuki-image/db"

func InitDataBase() error {
	err := db.InitDataBase()
	if err != nil {
		return err
	}
	return nil
}
