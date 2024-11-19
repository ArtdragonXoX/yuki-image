package album

import (
	"fmt"
	"time"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
	"yuki-image/internal/model"
	"yuki-image/utils"
)

func GetAllSize() (uint64, error) {
	size, err := utils.GetDirSize(conf.Conf.Image.Path)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func GetAllCount() (uint64, error) {
	count, err := utils.GetFileCount(conf.Conf.Image.Path)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetSize(id uint64) (uint64, error) {
	name, err := db.SelectAlbumNameFromId(id)
	if err != nil {
		return 0, err
	}
	return GetSizeFromName(name)
}

func GetCount(id uint64) (uint64, error) {
	name, err := db.SelectAlbumNameFromId(id)
	if err != nil {
		return 0, err
	}
	return GetCountFromName(name)
}

func GetSizeFromName(name string) (uint64, error) {
	path := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, name)
	size, err := utils.GetDirSize(path)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func GetCountFromName(name string) (uint64, error) {
	path := fmt.Sprintf("%s/%s", conf.Conf.Image.Path, name)
	count, err := utils.GetFileCount(path)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetCountStatistics(id uint64, dateS time.Time, dateE time.Time) (model.Statictics, error) {
	statistics, err := db.SelectCountStatistics(id, dateS, dateE)
	if err != nil {
		return nil, err
	}
	statistics.FillZero(dateS, dateE)
	return statistics, nil
}

func GetCountStatisticsFromName(name string, dateS time.Time, dateE time.Time) (model.Statictics, error) {
	id, err := db.SelectAlbumIdFromName(name)
	if err != nil {
		return nil, err
	}
	return GetCountStatistics(id, dateS, dateE)
}

func GetAllCountStatistics(dateS time.Time, dateE time.Time) (model.Statictics, error) {
	statictics, err := db.SelectAllCountStatistics(dateS, dateE)
	if err != nil {
		return nil, err
	}
	statictics.FillZero(dateS, dateE)
	return statictics, nil
}
