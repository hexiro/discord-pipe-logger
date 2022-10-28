package main

import (
	"github.com/hexiro/discord-pipe-logger/cli"
	"github.com/hexiro/discord-pipe-logger/pipe"
	"github.com/hexiro/discord-pipe-logger/webhook"
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
		_ = hook.SendMessage(&webhook.Message{
			Content: message,
		})
	}
}
