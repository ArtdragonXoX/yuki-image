package utils

import (
	"os"
	"path/filepath"
	"sync"
)

var fileCounter = make(map[uint64]uint64)
var fileMutexMap = make(map[uint64]*sync.Mutex)

func AddFileCounter(id uint64) {
	if _, ok := fileMutexMap[id]; !ok {
		fileMutexMap[id] = &sync.Mutex{} // 初始化新的互斥锁
	}
	fileMutexMap[id].Lock()
	defer fileMutexMap[id].Unlock()
	fileCounter[id]++
}

func SubFileCounter(id uint64) {
	if _, ok := fileMutexMap[id]; !ok {
		fileMutexMap[id] = &sync.Mutex{} // 初始化新的互斥锁
	}
	fileMutexMap[id].Lock()
	defer fileMutexMap[id].Unlock()
	fileCounter[id]--
}

func GetFileCounter(id uint64) uint64 {
	if _, ok := fileMutexMap[id]; !ok {
		fileMutexMap[id] = &sync.Mutex{} // 初始化新的互斥锁
	}
	fileMutexMap[id].Lock()
	defer fileMutexMap[id].Unlock()
	return fileCounter[id]
}

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
