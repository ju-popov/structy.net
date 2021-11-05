package decompressbraces

import (
	"strings"
)

type stackValue struct {
	count      int
	startIndex int
}

func Iterative(str string) string {
	stack := make([]stackValue, 0)

	for currentIndex := 0; currentIndex < len(str); currentIndex++ {
		char := str[currentIndex]

		if char == '{' {
			stack = append(stack, stackValue{count: int(str[currentIndex-1] - '0'), startIndex: currentIndex - 1})
		} else if (char == '}') && (len(stack) > 0) {
			startIndex := stack[len(stack)-1].startIndex

			left := str[0:startIndex]

			compressed := str[startIndex+2 : currentIndex]
			count := stack[len(stack)-1].count
			decompressed := strings.Repeat(compressed, count)

			right := str[currentIndex+1:]

			str = left + decompressed + right

			//nolint:gomnd
			currentIndex += len(decompressed) - (len(compressed) + 3)

			stack = stack[:len(stack)-1]
		}
	}

	return str
}
