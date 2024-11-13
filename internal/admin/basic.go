package admin

import (
	"errors"
	"log"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func UserLogin(u model.User) (string, error) {

	// 校验参数
	if u.Pwd == "" || u.Name == "" {
		return "", errors.New("参数错误，名称或密码不能为空")
	}
	var admin model.User

	err := utils.ReadYaml(&admin, "admin.yaml")
	if err != nil {
		log.Println(err)
		return "", errors.New("读取配置文件失败")
	}
	if err := admin.VerifyByPassword(u.Pwd); err != nil || u.Name != admin.Name {
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	//token := "{test token}"
	token, err := utils.GenerateToken(u.Name)
	if err != nil || token == "" {
		if err != nil {
			log.Println(err)
		}
		return "", errors.New("登录失败")
	}

	// 登录成功，返回token
	return token, nil
}

func UserRegister(u model.User) error {
	if ok, err := utils.IsFileExists("admin.yaml"); err != nil || ok {
		return errors.New("管理员已存在，请勿重复注册")
	}

	u.HashPassword()
	utils.WriteYaml(u, "admin.yaml")

	return nil
}
