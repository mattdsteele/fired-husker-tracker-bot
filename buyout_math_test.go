package main

import (
	"testing"
	"time"
)

func TestMonthMath(t *testing.T) {
	t.Log("Months Until")
	{
		may19 := MonthsUntil(2019, time.May)
		if may19 != 8 {
			t.Fatal("Could not do it")
		}
		oct18 := MonthsUntil(2018, time.October)
		if oct18 != 1 {
			t.Fatal("Failed math")
		}
	}
	t.Log("Pelini Months")
	{
		peliniMonths := PeliniMonths()
		if peliniMonths != 5 {
			t.Fatal("Bad Pelini")
		}
	}
	t.Log("Riley Months")
	{
		rileyMonths := RileyMonths()
		if rileyMonths != 29 {
			t.Fatal("Bad Riley")
		}
	}
	t.Log("Eichorst Months")
	{
		eichorstMonths := EichorstMonths()
		if eichorstMonths != 9 {
			t.Fatal(eichorstMonths)
		}
	}
}
