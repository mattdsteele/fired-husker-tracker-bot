package main

import "fmt"

var peliniPayout int
var eichorstPayout int
var rileyPayout int

func init() {
	peliniPayout = 128009
	eichorstPayout = 62666
	rileyPayout = 165833
}
func TweetPeliniPayout() string {
	peliniMonths := PeliniMonths()
	tweet := fmt.Sprintf("Bo Pelini has %d months left until Nebraska completes his contract buyout. They will pay him %d this month and next month until February 2019.", peliniMonths, peliniPayout)
	return tweet
}

func TweetPeliniLeft() string {
	peliniMonths := PeliniMonths()
	tweet := fmt.Sprintf("Nebraska still owes $%d to Bo Pelini before they are done buying out his contract.", peliniMonths*peliniPayout)
	return tweet
}

func TweetRileyPayout() string {
	rileyMonths := RileyMonths()
	tweet := fmt.Sprintf("Mike Riley has %d months left until Nebraska completes his contract buyout. They will pay him %d this month and next month until February 2021.", rileyMonths, rileyPayout)
	return tweet
}

func TweetRileyLeft() string {
	rileyMonths := RileyMonths()
	tweet := fmt.Sprintf("Nebraska still owes $%d to Mike Riley before they are done buying out his contract.", rileyMonths*rileyPayout)
	return tweet
}
func TweetEichorstLeft() string {
	eichorstMonths := EichorstMonths()
	tweet := fmt.Sprintf("Nebraska still owes $%d to Shawn Eichorst before they are done buying out his contract.", eichorstMonths*eichorstPayout)
	return tweet
}
func ThisMonthPayout() string {
	tweet := fmt.Sprintf("This month, Nebraska will pay $%d in total to Bo Pelini, Mike Riley, and Shawn Eichorst this month.", peliniPayout+eichorstPayout+rileyPayout)
	return tweet

}
