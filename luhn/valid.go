package luhn

import (
	"fmt"
	"strconv"
)

// Valid checks whether card number is valid
func Valid(card string) bool {
	// replace all spaces
	// replacedCard := strings.ReplaceAll(card, " ", "")
	fmt.Println("the function called with", card)
	number := []int64{}
	for _, e := range []byte(card) {
		num, err := strconv.ParseInt(string(e), 10, 64)
		if err != nil {
			if string(e) == " " {
				continue
			}
			return false
		}
		number = append(number, num)
	}
	if len(number) < 2 {
		return false
	}

	var sum int64
	leng := len(number)

	for i, e := range number {
		if leng%2 == 0 { // even
			if i%2 == 0 {
				modNum := (e * 2 % 9)
				sum += modNum
			} else {
				sum += e
			}
		} else { // odd
			if i%2 == 1 {
				modNum := (e * 2 % 9)
				fmt.Println("the length odd", modNum)
				sum += modNum
			} else {
				sum += e
				fmt.Println("the sum", sum)
			}
		}
	}

	if sum%10 != 0 {
		return false
	}
	// convert to array of numbers
	return true
}
