package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	fmt.Println(now)

	var utc = time.Date(2008, time.July, 12, 13, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())

}
