package raindrops

import (
	"strconv"
)

func primeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Convert converts number to string
func Convert(a int) string {
	factors := primeFactors(a)
	sl := []int{}
	// fmt.Printf("%v", factors)
	anyof357 := false
	returnString := ""
	for i := 0; i < len(factors); i++ {
		switch factors[i] {
		case 3:
			if contains(sl, 3) {
				break
			}
			returnString = "Pling"
			anyof357 = true
			sl = append(sl, 3)
			break
		case 5:
			if contains(sl, 5) {
				break
			}
			returnString += "Plang"
			anyof357 = true
			sl = append(sl, 5)
			break
		case 7:
			if contains(sl, 7) {
				break
			}
			returnString += "Plong"
			anyof357 = true
			sl = append(sl, 7)
			break
		}
	}
	if !anyof357 {
		return strconv.Itoa(a)
	}
	return returnString
}
