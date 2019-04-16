package grains

import "errors"

var m map[int]uint64

// Square returns square of a number
func Square(num int) (uint64, error) {
	m = make(map[int]uint64)

	if num < 1 || num > 64 {
		return 0, errors.New("Invalid num")
	}
	if num == 1 {
		return 1, nil
	}
	if num == 2 {
		return 2, nil
	}
	elem, ok := m[num]
	if ok {
		return elem, nil
	}
	inter, _ := Square(num - 1)
	// inter2, _ := Square(num - 2)
	m[num] = 2 * inter
	return 2 * (inter), nil
}

// Total returns 64*64
func Total() uint64 {
	var result uint64
	result = 0
	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		result += val
	}
	return result
}
