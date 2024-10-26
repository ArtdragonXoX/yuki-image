package utils

import "os"

func EnsureDir(dir string) error {
	// 检查文件夹是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 文件夹不存在，尝试创建文件夹
		err := os.MkdirAll(dir, 0755) // 0755是权限设置，表示所有者有读写执行权限，其他用户有读和执行权限
		if err != nil {
			return err // 如果创建失败，返回错误
		}
	}
	return nil // 文件夹存在或创建成功
}
