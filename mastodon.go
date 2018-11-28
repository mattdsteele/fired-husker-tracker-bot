package bot

import (
	"context"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

func Connect() *mastodon.Client {
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mastodon.social",
		ClientID:     os.Getenv("MASTODON_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTODON_CLIENT_SECRET"),
		AccessToken:  os.Getenv("MASTODON_ACCESS_TOKEN"),
	})
	return client
}

func Toot(client *mastodon.Client, toot string) {
	_, err := client.PostStatus(context.Background(), &mastodon.Toot{Status: toot, Visibility: "unlisted"})
	if err != nil {
		log.Fatal(err)
	}

}
