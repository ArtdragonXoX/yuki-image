package tmp

import (
	"yuki-image/internal/model"
	"yuki-image/utils"
)

var TmpPath string

func GetInfo() (model.TmpInfo, error) {
	var tmpInof model.TmpInfo
	size, err := GetSize()
	if err != nil {
		return tmpInof, err
	}
	tmpInof.Size = size
	count, err := GetCount()
	if err != nil {
		return tmpInof, err
	}
	tmpInof.Count = count
	return tmpInof, nil
}

func GetSize() (uint64, error) {
	return utils.GetDirSize(TmpPath)
}

func GetCount() (uint64, error) {
	return utils.GetFileCount(TmpPath)
}

func Clear() error {
	var err error
	err = utils.DeleteDir(TmpPath)
	_ = utils.EnsureDir(TmpPath)
	return err
}
