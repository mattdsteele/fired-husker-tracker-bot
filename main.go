package main
import (
	"fmt"
)

func mainMastodon() {

	connection := Connect()
	for true {
		Toot(connection, TweetRemainingMonths())
		fmt.Println("Tooted")
	}
}

func mainStandaloneTwitter() {
	tweetList := TweetRileyPayout()
	fmt.Println("Going to tweet:", tweetList)
	_, err := Tweet(tweetList)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

