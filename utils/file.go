package utils

import (
	"os"
	"path/filepath"
	"strings"
)

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

func GetDirSize(dir string) (uint64, error) {
	var size int64
	err := filepath.Walk(dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return uint64(size), err
}

func GetFileCount(path string) (uint64, error) {
	var count int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return uint64(count), err
}

func DeleteDir(dir string) error {
	return os.RemoveAll(dir)
}

func GetFileExt(filename string) string {
	// 使用 strings.LastIndex 查找最后一个点（.）的位置
	dotIndex := strings.LastIndex(filename, ".")
	if dotIndex == -1 {
		return ""
	}
	// 返回从最后一个点开始到文件名末尾的子串作为文件扩展名
	return filename[dotIndex+1:]
}
