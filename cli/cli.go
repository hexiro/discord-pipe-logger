package cli

import (
	"discord-pipe-logger/webhook"
	"flag"
	"os"
)

func Parse() *webhook.Webhook {
	var id string
	var token string
	flag.StringVar(&id, "i", "", "The ID of the Discord webhook")
	flag.StringVar(&token, "t", "", "The Token of the Discord webhook")
	flag.Parse()

	if id == "" || token == "" {
		flag.Usage()
		os.Exit(1)
	}

	hook, err := webhook.FromIDAndToken(id, token)
	if err != nil {
		panic(err)
	}
	return hook
}
