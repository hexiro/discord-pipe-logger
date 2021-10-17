package cli

import (
	"discord-pipe-logger/webhook"
	"errors"
	"os"
	"strings"
)

func Parse() (*webhook.Webhook, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("no input provided")
	}

	var hook *webhook.Webhook
	var err error

	arg := os.Args[1]
	hook, err = webhook.FromURL(arg)

	if err == nil {
		return hook, err
	}
	argSplit := strings.Split(arg, "/")
	argSplitLength := len(argSplit)

	if argSplitLength < 2 {
		return  hook, errors.New("hook can't be parsed from input")
	}

	id, token := argSplit[argSplitLength-2], argSplit[argSplitLength-1]
	return webhook.FromIDAndToken(id, token)
}
