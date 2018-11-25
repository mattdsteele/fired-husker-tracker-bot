package main
import (
	"fmt"
	"math/rand"
	"time"

	"github.com/leekchan/accounting"
)

var peliniPayout int
var eichorstPayout int
var rileyPayout int
var ac accounting.Accounting

func init() {
	peliniPayout = 128009
	eichorstPayout = 62666
	rileyPayout = 165833
	ac = accounting.Accounting{Symbol: "$", Precision: 0}
}

func dollar(money int) string {
	return ac.FormatMoney(money)
}

func TweetPeliniPayout() string {
	peliniMonths := PeliniMonths()
	tweet := fmt.Sprintf("Bo Pelini has %d months left until Nebraska completes his contract buyout. They will pay him %s this month and next month until February 2019.", peliniMonths, dollar(peliniPayout))
	return tweet
}

func TweetPeliniLeft() string {
	peliniMonths := PeliniMonths()
	tweet := fmt.Sprintf("Nebraska still owes %s to Bo Pelini before they are done buying out his contract.", dollar(peliniMonths*peliniPayout))
	return tweet
}

func TweetRileyPayout() string {
	rileyMonths := RileyMonths()
	tweet := fmt.Sprintf("Mike Riley has %d months left until Nebraska completes his contract buyout. They will pay him %s this month and next month until February 2021.", rileyMonths, dollar(rileyPayout))
	return tweet
}

func TweetRileyLeft() string {
	rileyMonths := RileyMonths()
	tweet := fmt.Sprintf("Nebraska still owes %s to Mike Riley before they are done buying out his contract.", dollar(rileyMonths*rileyPayout))
	return tweet
}
func TweetEichorstLeft() string {
	eichorstMonths := EichorstMonths()
	tweet := fmt.Sprintf("Nebraska still owes %s to Shawn Eichorst before they are done buying out his contract.", dollar(eichorstMonths*eichorstPayout))
	return tweet
}
func ThisMonthPayout() string {
	tweet := fmt.Sprintf("This month, Nebraska will pay %s in total to Bo Pelini, Mike Riley, and Shawn Eichorst.", dollar(peliniPayout+eichorstPayout+rileyPayout))
	return tweet

}

func RandomTweet() string {
	rand.Seed(time.Now().UnixNano())
	toots := []string{
		ThisMonthPayout(),
		TweetEichorstLeft(),
		TweetRileyLeft(),
		TweetRileyPayout(),
		TweetPeliniLeft(),
		TweetPeliniPayout(),
	}
	return toots[rand.Intn(len(toots))]
}

