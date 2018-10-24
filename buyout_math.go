package main

import (
	"time"
)

func MonthsUntil(year int, month time.Month) int {
	now := time.Now()
	yearDelta := year - now.Year()
	nowAsInt := int(now.Month())
	monthDelta := nowAsInt - int(month)
	return (yearDelta * 12) - monthDelta + 1
}

func PeliniMonths() int {
	return MonthsUntil(2019, time.February)
}
func RileyMonths() int {
	return MonthsUntil(2021, time.February)
}
func EichorstMonths() int {
	return MonthsUntil(2019, time.June)
}
