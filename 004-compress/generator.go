package compress

import (
	"fmt"
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
		var (
			count    int
			lastElem rune
		)

		for _, elem := range s {
			if elem == lastElem {
				count++
			} else {
				if count > 0 {
					c <- struct {
						number int
						char   string
					}{number: count, char: string(lastElem)}
				}

				lastElem = elem
				count = 1
			}
		}

		if count > 0 {
			c <- struct {
				number int
				char   string
			}{number: count, char: string(lastElem)}
		}

		close(c)
	}()

	return c
}

func Generator(s string) string {
	parts := make([]string, 0)

	for match := range findAllStringSubmatchGenerator(s) {
		if match.number == 1 {
			parts = append(parts, match.char)
		} else if match.number > 1 {
			parts = append(parts, fmt.Sprintf("%d%s", match.number, match.char))
		}
	}

	return strings.Join(parts, "")
}
