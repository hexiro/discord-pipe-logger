package pipe

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// some code taken from
// https://zetcode.com/golang/pipe/

func ReadInput() []byte {
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {

		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Hello %s!\n", buf)
		return buf
	}
	return nil
}
