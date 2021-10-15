package pipe

import (
	"bufio"
	"log"
	"math"
	"os"
)

// some code taken from
// https://zetcode.com/golang/pipe/

const discordCharLimit = 2000

func ReadMessages() []string {
	stat, err := os.Stdin.Stat()
	var messages []string

	if err != nil {
		return messages
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {

		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}


		str := string(buf)

		dividedLength := float64(len(str)) / float64(discordCharLimit)
		messagedNeeded := int(math.Ceil(dividedLength))

		for i := 0; i < messagedNeeded; i++ {
			lowerBoundary := i*discordCharLimit
			upperBoundary := (i+1)*discordCharLimit
			if upperBoundary > len(str) {
				upperBoundary = len(str)
			}
			messages = append(messages, str[lowerBoundary:upperBoundary])
		}

		return messages
	}
	return nil
}
