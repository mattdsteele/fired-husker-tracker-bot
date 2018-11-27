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
	eichorstPayout = 762439
	rileyPayout = 6235393
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
	tweet := fmt.Sprintf("Mike Riley was paid %s in a lump sum in December 2017 to fulfil Nebraska's contract obligation. https://journalstar.com/sports/huskers/football/nu-paid-former-coach-riley-lump-sum-of-more-than/article_f0b24796-b2a6-50ec-a68f-16ef42cdc179.amp.html", dollar(rileyPayout))
	return tweet
}

func TweetEichorstLeft() string {
	tweet := fmt.Sprintf("Shawn Eichorst was paid %s in a lump sum in June 2018 to buy out the remainder of his contract.", dollar(eichorstPayout))
	return tweet
}
func ThisMonthPayout() string {
	tweet := fmt.Sprintf("This month, Nebraska will pay %s to Bo Pelini as they continue to work off his contractual obligations.", dollar(peliniPayout))
	return tweet
}

func TweetRemainingMonths() string {
	return fmt.Sprintf("There are %d months left until Scott Frost is the only head coach drawing a salary from the University of Nebraska.", MonthsUntil(2019, time.March))
}

func RandomTweet() string {
	rand.Seed(time.Now().UnixNano())
	toots := []string{
		ThisMonthPayout(),
		TweetEichorstLeft(),
		TweetRileyPayout(),
		TweetPeliniLeft(),
		TweetPeliniPayout(),
		TweetRemainingMonths(),
	}
	return toots[rand.Intn(len(toots))]
}

