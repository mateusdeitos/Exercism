package secret

import (
	"slices"
	"strconv"
)

const (
	wink          = "wink"
	doubleBlink   = "double blink"
	closeYourEyes = "close your eyes"
	jump          = "jump"
)

type modifier struct {
	index  int
	action string
	modify func(binary string, c []string) []string
}

func (m *modifier) apply(binary string, c []string) []string {
	if len(binary) < m.index {
		return c
	}

	if binary[len(binary)-m.index] == '1' {
		if m.modify != nil {
			c = m.modify(binary, c)
		} else {
			c = append(c, m.action)
		}
	}

	return c
}

func Handshake(code uint) []string {
	binary := strconv.FormatUint(uint64(code), 2)

	m := []modifier{
		{1, wink, nil},
		{2, doubleBlink, nil},
		{3, closeYourEyes, nil},
		{4, jump, nil},
		{5, "", func(binary string, c []string) []string {
			slices.Reverse(c)
			return c
		}},
	}

	var c []string
	for i := 0; i < len(m); i++ {
		c = m[i].apply(binary, c)
	}

	return []string(c)
}
