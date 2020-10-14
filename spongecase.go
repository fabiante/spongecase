package spongecase

import (
	"math/rand"
	"unicode"
)

func ApplyStr(in string) string {
	out := make([]rune, len(in))
	var newChar rune

	for pos, char := range []rune(in) {
		toUpper := rand.Float32() > 0.5
		if toUpper {
			newChar = unicode.ToUpper(char)
		} else {
			newChar = unicode.ToLower(char)
		}
		out[pos] = newChar
	}

	return string(out)
}
