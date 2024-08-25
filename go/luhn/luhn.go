package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	if len(strings.Trim(id, " ")) <= 1 {
		return false
	}

	i := len(id) - 1
	sum := 0
	double := false
	for i >= 0 {
		if id[i] == ' ' {
			i--
			continue
		}

		v, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if double {
			v = v * 2
			if v > 9 {
				v = v - 9
			}
		}

		double = !double
		sum += v
		i--
	}

	return sum%10 == 0
}
