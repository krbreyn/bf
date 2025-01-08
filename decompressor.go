package bf

import (
	"strconv"
	"strings"
	"unicode"
)

func DecompressProgram(input string) string {
	if len(input) == 0 {
		return ""
	}

	sb := strings.Builder{}

	var currRune rune

	for i, c := range input {
		if !unicode.IsDigit(c) {
			sb.WriteRune(c)
			currRune = c
		} else {
			num := string(c)

			next := input[i+1]
			for unicode.IsDigit(rune(next)) {
				i++
				num += string(input[i])
				next = input[i+1]
			}

			count, _ := strconv.Atoi(num)
			for range count - 1 {
				sb.WriteRune(currRune)
			}
		}
	}

	return sb.String()
}
