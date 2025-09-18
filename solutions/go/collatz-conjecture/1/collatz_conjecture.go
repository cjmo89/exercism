package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid n")
	}
	counter := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			counter++
		} else {
			n = 3*n + 1
			counter++
		}
	}
	return counter, nil
}
