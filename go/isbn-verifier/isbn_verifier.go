package isbn

import (
	"strconv"
)

func IsValidISBN(isbn string) bool {
	if len(isbn) != 13 && len(isbn) != 10 {
		return false
	}

	total := 0
	count := 0
	for i, c := range isbn {
		isLast := i == len(isbn)-1
		if string(c) == "-" {
			continue
		}

		cStr := string(c)
		if isLast && string(c) == "X" {
			cStr = "10"
		}

		v, err := strconv.Atoi(cStr)
		if err != nil {
			return false
		}

		total += v * (10 - count)
		count++

	}

	return total%11 == 0
}
