package robotname

import "math/rand"

// Robot definiton
type Robot struct {
	name string
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("1234567890")
var names = map[string]int{}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		if i < 2 {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		} else {
			b[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	}
	return string(b)
}

// Name should generate new name
func (r *Robot) Name() (string, error) {
	if len(r.name) < 1 {
		name := randStringRunes(5)
		for {
			_, chk := names[name]
			if chk {
				name = randStringRunes(5)
			} else {
				break
			}
		}
		r.name = name
		names[r.name] = 1
	}

	return r.name, nil
}

// Reset should rest the name
func (r *Robot) Reset() {
	name := randStringRunes(5)
	for {
		_, chk := names[name]
		if chk {
			name = randStringRunes(5)
		} else {
			break
		}
	}
	r.name = name
	names[r.name] = 1
}
