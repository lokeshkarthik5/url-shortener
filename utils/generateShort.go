package utils

import (
	"math/rand/v2"
	"strings"
)

func GenerateShortUrls() string {
	rand.IntN(100000)
	var sb strings.Builder

	for i := 0; i < 4; i++ {
		char := 'A' + rune(rand.IntN(26))
		sb.WriteRune(char)
	}
	for i := 0; i < 4; i++ {
		digit := '0' + rune(rand.IntN(10))
		sb.WriteRune(digit)
	}

	return sb.String()
}
