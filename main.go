package main

import (
	"discord-pipe-logger/cli"
	"discord-pipe-logger/pipe"
	"discord-pipe-logger/webhook"
)

func printUsage(err error) {
	println("Error:", err.Error())
	println("Usage: {command} | discord-pipe-logger \"{webhook}\"")
}

func main() {
	messages, err := pipe.ReadMessages()
	if err != nil {
		printUsage(err)
		return
	}
	hook, err := cli.Parse()
	if err != nil {
		printUsage(err)
		return
	}
	for _, message := range messages {
		_ = hook.SendMessage(&webhook.WebhookMessage{
			Content: message,
		})
	}
}
