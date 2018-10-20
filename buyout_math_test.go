package main

import (
	"testing"
	"time"
)

func TestMonthMath(t *testing.T) {
	t.Log("Months Until")
	{
		may19 := MonthsUntil(2019, time.May)
		if may19 != 7 {
			t.Fatal("Could not do it")
		}
	}
}
