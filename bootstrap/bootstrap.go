package bootstrap

func InitServer() {
	InitConfig()
	InitFile()
	InitDataBase()
	InitHandlers()
}
