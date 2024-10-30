package tmp

import (
	"yuki-image/internal/model"
	"yuki-image/utils"
)

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
	return utils.GetDirSize("tmp")
}

func GetCount() (uint64, error) {
	return utils.GetFileCount("tmp")
}

func Clear() error {
	var err error
	err = utils.DeleteDir("tmp")
	_ = utils.EnsureDir("tmp")
	return err
}
