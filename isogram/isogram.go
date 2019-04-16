package isogram

import "strings"

// IsIsogram ruturns whether the word is isogram or not
func IsIsogram(word string) bool {
	alphaMap := map[byte]int{}
	word = strings.ToLower(word)
	for _, e := range []byte(word) {
		if string(e) == " " || string(e) == ("-") {
			continue
		}
		_, ok := alphaMap[e]
		if ok {
			return false
		}
		alphaMap[e] = 1
	}
	return true
}
