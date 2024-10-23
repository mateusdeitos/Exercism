package raindrops

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	var sounds string
	if number%3 == 0 {
		sounds = fmt.Sprintf("%s%s", sounds, "Pling")
	}

	if number%5 == 0 {
		sounds = fmt.Sprintf("%s%s", sounds, "Plang")
	}

	if number%7 == 0 {
		sounds = fmt.Sprintf("%s%s", sounds, "Plong")
	}

	if sounds != "" {
		return sounds
	}

	return strconv.Itoa(number)
}
