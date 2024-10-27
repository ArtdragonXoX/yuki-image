package bootstrap

import (
	"yuki-image/server"
)

func InitServer() error {
	return server.NewAndInit()
}
