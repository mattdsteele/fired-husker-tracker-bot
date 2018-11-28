package main
import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	bot "github.com/mattdsteele/husker-buyout-bot"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.CloudWatchEvent) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.Source)
	connection := bot.Connect()
	tweet := bot.RandomTweet()
	bot.Toot(connection, tweet)
	log.Printf("Tooted: %s\n", tweet)
}
func main() {
	lambda.Start(Handler)
}

