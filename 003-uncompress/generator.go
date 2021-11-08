package uncompress

import (
	"strings"
)

//nolint:varnamelen
func findAllStringSubmatchGenerator(s string) chan struct {
	number int
	char   string
} {
	//nolint:varnamelen
	c := make(chan struct {
		number int
		char   string
	})

	go func() {
		number := 0

		for _, elem := range s {
			if (elem >= '0') && (elem <= '9') {
				//nolint:gomnd
				number = number*10 + int(elem-'0')
			} else {
				c <- struct {
					number int
					char   string
				}{number: number, char: string(elem)}
				number = 0
			}
		}

		close(c)
	}()

	return c
}

func Generator(s string) string {
	parts := make([]string, 0)

	for match := range findAllStringSubmatchGenerator(s) {
		parts = append(parts, strings.Repeat(match.char, match.number))
	}

	return strings.Join(parts, "")
}
