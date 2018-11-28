package main
import (
	"fmt"

	bot "github.com/mattdsteele/husker-buyout-bot"
)

func main() {

	connection := bot.Connect()
	bot.Toot(connection, bot.TweetRemainingMonths())
	fmt.Println("Tooted")
}

