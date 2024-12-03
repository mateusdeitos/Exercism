package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

const minutesInADay = 24 * 60

func New(h, m int) Clock {
	h, m = normalize(h, m)
	return Clock{h: h, m: m}
}

func (c Clock) Add(m int) Clock {
	hours, min := normalize(c.h, c.m+m)
	c.h = hours
	c.m = min

	return c
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// normalize normalizes hours and minutes into the range [0, 24) and [0, 60)
// negative values are allowed and
// values greater than 24 (for hours) and 60 (for minutes) are reduced to the equivalent value in [0, 24) and [0, 60)
// ex:
// normalize(23, 61) = (0, 1)
// normalize(-1, 1) = (23, 1)
// normalize(50, 125) = (4, 5)
func normalize(h, m int) (int, int) {
	h = h % 24
	if h < 0 {
		h += 24
	}

	minutes := m % minutesInADay
	if minutes < 0 {
		minutes += minutesInADay
	}

	minutesToHours := int(minutes / 60)
	h += minutesToHours
	h = h % 24

	m = minutes - minutesToHours*60
	m = m % 60
	return h, m
}
