package hamming

import "errors"
import "strings"

/*
Distance function takes two string arguments and gives the diff count
*/
func Distance(a, b string) (int, error) {
	aLength := len(a)
	bLength := len(b)
	if aLength != bLength {
		return -1, errors.New("Invalid")
	}
	if aLength == 0 && bLength == 0 {
		return 0, nil
	}
	aChars := strings.Split(a, "")
	bChars := strings.Split(b, "")
	count := 0
	for i, v := range aChars {
		if v != bChars[i] {
			count++
		}
	}
	return count, nil
}
