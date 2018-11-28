package main
import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	bot "github.com/mattdsteele/husker-buyout-bot"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func TwitterHandler(request events.CloudWatchEvent) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.Source)
	tweet := bot.RandomTweet()
	bot.Tweet(tweet)
	log.Printf("Tooted: %s\n", tweet)
}
func main() {
	lambda.Start(TwitterHandler)
}

