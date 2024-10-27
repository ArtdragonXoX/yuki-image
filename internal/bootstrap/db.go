package bootstrap

import "yuki-image/internal/db"

func InitDataBase() error {
	return db.InitDataBase()
}
