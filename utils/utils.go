package utils

import (
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

var rc *redis.Client
var layout string = "2006-01-02 15:04:05"

func GetRedisClient() *redis.Client {

	if rc != nil {
		return rc
	}

	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetCurDate() string {
	return time.Now().Format(layout)
}

func GetCurTs() int64 {
	return time.Now().Unix()
}

func GetTsToDate(ts int64) string {
	return time.Unix(ts, 0).Format(layout)
}

func GetDateToTs(date string) int64 {
	ts, _ := time.ParseInLocation(layout, date, time.Local)
	return ts.Unix()
}

func GenRandStrings(max int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"

	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < max; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetHotPlateMoment() map[string]bool {
	m := make(map[string]bool)
	m["12:01"] = true
	m["15:01"] = true
	m["12:35"] = true
	return m
}