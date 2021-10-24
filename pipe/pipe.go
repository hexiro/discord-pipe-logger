package pipe

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// some code taken from
// https://zetcode.com/golang/pipe/

const discordCharLimit = 2000

func Read() ([]rune, error) {
	info, err := os.Stdin.Stat()
	var output []rune

	if err != nil {
		return output, err
	}

	if (info.Mode() & os.ModeCharDevice) != 0 {
		return output, errors.New("failed to read from pipe")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return output, nil
}

func ReadMessages() ([]string, error) {
	var messages []string

	input, err := Read()

	if err != nil {
		return messages, err
	}

	inputString := string(input)
	// split every discordCharLimit (2000) chars

	for i := 0; i < len(inputString); i += discordCharLimit {
		lowerBoundary := i
		upperBoundary := i + discordCharLimit
		if upperBoundary > len(inputString) {
			upperBoundary = len(inputString)
		}
		messages = append(messages, inputString[lowerBoundary:upperBoundary])
	}
	return messages, nil
}
