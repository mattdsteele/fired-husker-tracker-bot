package main
import (
	"fmt"

	bot "github.com/mattdsteele/husker-buyout-bot"
)

func main() {
	tweetList := bot.TweetRileyPayout()
	fmt.Println("Going to tweet:", tweetList)
	_, err := bot.Tweet(tweetList)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

