package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}

	steps := 0
	for {
		if n == 1 {
			break
		}

		steps++
		if isEven(n) {
			n = onEven(n)
			continue
		}

		n = onOdd(n)
	}

	return steps, nil
}

func isEven(n int) bool {
	return n%2 == 0
}

func onEven(n int) int {
	return int(n / 2)
}

func onOdd(n int) int {
	return 3*n + 1
}
