package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/exp/rand"
)

// var timeMutex = &sync.Mutex{}
var KeyLength *int
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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
	// timeMutex.Lock()
	// defer timeMutex.Unlock()
	// time.Sleep(time.Microsecond)

	timestamp := time.Now().UnixNano()
	timestampBytes := []byte(fmt.Sprintf("%d", timestamp))

	dataTOHash := append(buff, timestampBytes...)
	// dataTOHash := append(buff, []byte(GetRandKey())...)
	hash := md5.Sum(dataTOHash)
	hashHex := fmt.Sprintf("%x", hash)
	return hashHex, nil
}

func GetRandKey() string {
	// timeMutex.Lock()
	// defer timeMutex.Unlock()
	// time.Sleep(time.Microsecond)
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]rune, *KeyLength)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

func WaitTcp(ip string, port string) {
	address := fmt.Sprintf("%s:%s", ip, port)
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", address)
		if err == nil {
			break
		}
		log.Println(fmt.Sprintf("Waiting for %s connection...", address))
		time.Sleep(time.Second)
	}
	defer conn.Close()
	log.Println(fmt.Sprintf("Connected to %s", address))
}

func GenerateDateList(startDate time.Time, endDate time.Time) []string {
	var dateList []string
	for startDate.Before(endDate.AddDate(0, 0, 1)) {
		dateList = append(dateList, startDate.Format("2006-01-02"))
		startDate = startDate.AddDate(0, 0, 1)
	}
	return dateList
}
