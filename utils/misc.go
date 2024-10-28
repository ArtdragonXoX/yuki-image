package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func Contains[T comparable](arr []T, value T) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func GetByteHash(buff []byte) (string, error) {
	timestamp := time.Now().UnixNano()
	timestampBytes := []byte(fmt.Sprintf("%d", timestamp))

	dataTOHash := append(buff, timestampBytes...)
	hash := md5.Sum(dataTOHash)
	hashHex := fmt.Sprintf("%x", hash)
	return hashHex, nil
}
