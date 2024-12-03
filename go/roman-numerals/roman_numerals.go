package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid input")
	}

	conversion := []struct {
		arabic int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	r := ""

	cursor := 0
	for input > 0 {
		c := conversion[cursor]

		if input >= c.arabic {
			input -= c.arabic
			r += c.roman
			continue
		}

		cursor++
	}

	return r, nil
}
