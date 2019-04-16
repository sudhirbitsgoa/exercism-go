package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var lock = &sync.Mutex{}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency frequency calculator
func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	done := make(chan bool, len(s))
	for i := 0; i < len(s); i++ {
		str := s[i]
		go func(st string) {
			for _, r := range st {
				lock.Lock()
				m[r]++
				lock.Unlock()
			}
			done <- false
		}(str)
	}
	for i := 0; i < len(s); i++ {
		<-done
	}
	lock.Lock()
	// fmt.Printf("%v", m)
	lock.Unlock()
	return m
}
