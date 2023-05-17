package util

import (
	"fmt"
	"time"
)

var countryTz = map[string]string{
	"Bangkok": "Asia/Bangkok",
}

func GetLocalTime(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func Time(s string) time.Time {
	fmt.Println(s)
	time, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println(err)
	}
	return time
}
