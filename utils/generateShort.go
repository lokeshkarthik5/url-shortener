package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateShortUrls() string {
	rand.Seed(time.Now().Unix())
	var sb strings.Builder

	for i := 0; i < 4; i++ {
		char := 'A' + rune(rand.Intn(26))
		sb.WriteRune(char)
	}
	for i := 0; i < 4; i++ {
		digit := '0' + rune(rand.Intn(10))
		sb.WriteRune(digit)
	}

	return sb.String()
}
