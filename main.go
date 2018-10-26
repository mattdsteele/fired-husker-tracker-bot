package main

import (
	"fmt"
)

/*
func mainMastodon() {

	connection := Connect()
	for true {
		Toot(connection, RandomTweet())
		fmt.Println("Tooted")
		time.Sleep(12 * time.Hour)
	}
}
*/

func mainTest() {
	tweetList := RandomTweet()
	tweet, err := Tweet(tweetList)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	fmt.Printf("Tweeted: %s", tweet)
}
