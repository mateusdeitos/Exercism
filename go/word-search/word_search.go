package wordsearch

import (
	"errors"
	"fmt"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, word := range words {
		found := false
		if err := searchHorizontally(result, word, puzzle); err == nil {
			found = true
		}

		if !found {
			if err := searchVertically(result, word, puzzle); err == nil {
				found = true
			}
		}

		if !found {
			if err := searchDiagonally(result, word, puzzle); err == nil {
				found = true
			}
		}

		if !found {
			result[word] = notFound()
			err = fmt.Errorf("%s not found", word)
		}
	}

	return result, err
}

func notFound() [2][2]int {
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func searchHorizontally(r map[string][2][2]int, word string, puzzle []string) error {
	yBoundary := len(puzzle) - 1
	xBoundary := len(puzzle[0]) - len(word)

	for y := 0; y <= yBoundary; y++ {
		for x := 0; x <= xBoundary; x++ {
			w := buildWord(x, y, Neutral, Positive, word, puzzle)
			y2 := y
			x2 := x + len(word) - 1

			// left to right
			if w == word {
				r[word] = [2][2]int{{x, y}, {x2, y2}}
				return nil
			}

			// right to left
			rw := reverseString(w)
			if rw == word {
				r[word] = [2][2]int{{x2, y2}, {x, y}}
				return nil
			}
		}
	}

	return errors.New("word not found")
}

func searchVertically(r map[string][2][2]int, word string, puzzle []string) error {
	yBoundary := len(puzzle) - len(word)
	xBoundary := len(puzzle[0])

	for x := 0; x < xBoundary; x++ {
		for y := 0; y <= yBoundary; y++ {
			w := buildWord(x, y, Positive, Neutral, word, puzzle)
			y2 := y + len(word) - 1
			x2 := x

			// top to bottom
			if w == word {
				r[word] = [2][2]int{{x, y}, {x2, y2}}
				return nil
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				r[word] = [2][2]int{{x2, y2}, {x, y}}
				return nil
			}
		}
	}

	return errors.New("word not found")
}

func searchDiagonally(r map[string][2][2]int, word string, puzzle []string) error {
	yBoundary := len(puzzle) - len(word)
	xBoundary := len(puzzle[0]) - len(word)

	// top to bottom
	for x := 0; x <= xBoundary; x++ {
		for y := 0; y <= yBoundary; y++ {
			w := buildWord(x, y, Positive, Positive, word, puzzle)
			y2 := y + len(word) - 1
			x2 := x + len(word) - 1

			// top to bottom
			if w == word {
				r[word] = [2][2]int{{x, y}, {x2, y2}}
				return nil
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				r[word] = [2][2]int{{x2, y2}, {x, y}}
				return nil
			}
		}
	}

	yBoundary = len(puzzle) - yBoundary
	// bottom to top
	for x := 0; x <= xBoundary; x++ {
		for y := len(puzzle) - 1; y >= yBoundary; y-- {
			w := buildWord(x, y, Negative, Positive, word, puzzle)
			y2 := y - len(word) + 1
			x2 := x + len(word) - 1

			// top to bottom
			if w == word {
				r[word] = [2][2]int{{x, y}, {x2, y2}}
				return nil
			}

			// bottom to top
			rw := reverseString(w)
			if rw == word {
				r[word] = [2][2]int{{x2, y2}, {x, y}}
				return nil
			}
		}
	}

	return errors.New("word not found")
}

const (
	Positive = 1 << iota
	Negative
	Neutral
)

func buildWord(
	col,
	row,
	rowDirection,
	colDirection int,
	word string,
	puzzle []string,
) string {
	y := row
	x := col

	if rowDirection == Neutral && colDirection == Neutral {
		panic("invalid direction")
	}

	w := ""

	for {
		if len(w) == len(word) {
			break
		}

		w += string(puzzle[y][x])
		if rowDirection == Positive {
			y++
		} else if rowDirection == Negative {
			y--
		}

		if colDirection == Positive {
			x++
		} else if colDirection == Negative {
			x--
		}
	}

	return w
}

func reverseString(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}

	return r
}
