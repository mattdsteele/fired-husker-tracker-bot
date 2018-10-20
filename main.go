package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello %s", "world!\n")
	now := time.Now()
	fmt.Println(MonthsUntil(2019, time.February))
	fmt.Printf(now.String())
}

func MonthsUntil(year int, month time.Month) int {
	now := time.Now()
	yearDelta := year - now.Year()
	nowAsInt := int(now.Month())
	monthDelta := nowAsInt - int(month)
	return (yearDelta * 12) - monthDelta
}
