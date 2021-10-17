package main

import (
	"discord-pipe-logger/cli"
	"discord-pipe-logger/pipe"
	"discord-pipe-logger/webhook"
	"log"
)

func main() {
	messages, err := pipe.ReadMessages()
	if err != nil {
		log.Fatal(err)
	}
	hook, err := cli.Parse()
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		_ = hook.SendMessage(&webhook.WebhookMessage{
			Content: message,
		})
	}
}
