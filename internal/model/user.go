package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name string `json:"name,omitempty" yaml:"name"`
	Pwd  string `json:"pwd,omitempty" yaml:"pwd"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Pwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Pwd = string(hashedPassword)
	return nil
}

// 验证密码
func (u *User) VerifyByPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(pw))
}
