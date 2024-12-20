package album

import (
	"fmt"
	"time"
	"yuki-image/internal/conf"
	"yuki-image/internal/db"
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

func GetStatistics(id uint64, dateS time.Time, dateE time.Time) (map[string]uint64, error) {
	statistics, err := db.SelectStatistics(id, dateS, dateE)
	if err != nil {
		return nil, err
	}
	return statistics, nil
}

func GetStatisticsFromName(name string, dateS time.Time, dateE time.Time) (map[string]uint64, error) {
	id, err := db.SelectAlbumIdFromName(name)
	if err != nil {
		return nil, err
	}
	return GetStatistics(id, dateS, dateE)
}

func GetAllStatistics(dateS time.Time, dateE time.Time) (map[string]uint64, error) {
	return db.SelectAllStatistics(dateS, dateE)
}
