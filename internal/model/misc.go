package model

import (
	"encoding/json"
	"time"
)

const (
	jsonDateLayout = "2006-01-02" // 设置你期望的日期格式
)

// CustomTime 是一个包装了 time.Time 的类型，它实现了 json.Unmarshaler 接口
type CustomTime struct {
	time.Time
}

// UnmarshalJSON 是 json.Unmarshaler 接口的实现
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	err := ct.FromString(s)
	if err != nil {
		return err
	}
	return nil
}

func (ct *CustomTime) FromString(s string) error {
	t, err := time.Parse(jsonDateLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct *CustomTime) Now() {
	ct.Time = time.Now()
}
