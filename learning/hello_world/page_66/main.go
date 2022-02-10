package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var year int = now.Year()
	fmt.Println("Year = ", year)

	var month int = int(now.Month())
	fmt.Println("Month = ", month)

	var day int = now.Day()
	fmt.Println("Day = ", day)
}
