package bootstrap

func Init() error {

	err := InitConfig()
	if err != nil {
		return err
	}
	err = InitFile()
	if err != nil {
		return err
	}
	err = InitDataBase()
	if err != nil {
		return err
	}
	err = InitServer()
	if err != nil {
		return err
	}
	err = InitMisc()
	if err != nil {
		return err
	}
	return nil
}
