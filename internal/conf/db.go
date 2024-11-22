package conf

import "errors"

type DatabaseConf struct {
	MaxConn int `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdle int `yaml:"max_idle_conns" json:"max_idle_conns"`
}

func (d *DatabaseConf) Default() {
	d.MaxConn = 10
	d.MaxIdle = 5
}

func (d *DatabaseConf) Update(new DatabaseConf) error {
	if err := new.Check(); err != nil {
		return err
	}
	d.MaxConn = new.MaxConn
	d.MaxIdle = new.MaxIdle
	return nil
}

func (d *DatabaseConf) Check() error {
	if d.MaxConn <= 0 || d.MaxIdle < 0 {
		return errors.New("invalid database config")
	}
	return nil
}
