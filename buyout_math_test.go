package bot
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
		nov18 := MonthsUntil(2018, time.November)
		if nov18 != 1 {
			t.Fatal("Failed math")
		}
	}
	t.Log("Pelini Months")
	{
		peliniMonths := PeliniMonths()
		if peliniMonths != 4 {
			t.Fatal("Bad Pelini")
		}
	}
	t.Log("Riley Months")
	{
		rileyMonths := RileyMonths()
		if rileyMonths != 28 {
			t.Fatal("Bad Riley")
		}
	}
	t.Log("Eichorst Months")
	{
		eichorstMonths := EichorstMonths()
		if eichorstMonths != 8 {
			t.Fatal(eichorstMonths)
		}
	}
}

