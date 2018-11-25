package main
import (
	"strings"
	"testing"
)

func TestTweets(t *testing.T) {
	t.Log("Tweets")
	{
		tweet := TweetPeliniPayout()
		if !strings.Contains(tweet, "Bo Pelini has") {
			t.Fatal("Failed tweet " + tweet)
		}
	}
	t.Log("Tweet 2")
	{
		tweet := TweetPeliniLeft()
		if !strings.Contains(tweet, "Nebraska still owes $") {
			t.Fatal("Failed tweet " + tweet)
		}
	}
	t.Log("Payout")
	{
		tweet := ThisMonthPayout()
		if !strings.Contains(tweet, "$356,508") {
			t.Fatal("Failed tweet " + tweet)
		}
	}
}

