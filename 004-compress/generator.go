package compress

import (
	"strconv"
	"strings"
)

//nolint:varnamelen
func findAllStringSubmatchGenerator(s string) chan struct {
	number int
	char   int32
} {
	//nolint:varnamelen
	c := make(chan struct {
		number int
		char   int32
	})

	go func() {
		var (
			count int
			char  int32
		)

		for index, elem := range s {
			if index == 0 {
				char = elem
				count = 1
			} else {
				if elem == char {
					count++
				} else {
					c <- struct {
						number int
						char   int32
					}{number: count, char: char}
					char = elem
					count = 1
				}
			}
		}

		c <- struct {
			number int
			char   int32
		}{number: count, char: char}

		close(c)
	}()

	return c
}

func Generator(s string) (string, error) {
	var result strings.Builder

	for match := range findAllStringSubmatchGenerator(s) {
		if match.number == 1 {
			result.WriteRune(match.char)
		} else if match.number > 1 {
			//nolint:gomnd
			result.WriteString(strconv.FormatInt(int64(match.number), 10))
			result.WriteRune(match.char)
		}
	}

	return result.String(), nil
}
