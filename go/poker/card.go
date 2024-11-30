package poker

import (
	"fmt"
	"regexp"
)

type card struct {
	raw  string
	rank int
	suit suit
}

var cardStrPattern = regexp.MustCompile(`^(1|2|3|4|5|6|7|8|9|10|J|Q|K|A){1}(♢|♤|♡|♧){1}$`)

func NewCard(s string) (*card, error) {
	conversion := map[string]int{
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
		"A":  ACE_UPPER_BOUND,
	}

	c := &card{}
	c.raw = s

	matches := cardStrPattern.FindStringSubmatch(s)
	if len(matches) == 3 {
		rank, ok := conversion[matches[1]]
		if !ok {
			return nil, fmt.Errorf("invalid card rank")
		}

		c.rank = rank
		c.suit = suit(matches[2])
		if !c.suit.isValid() {
			return nil, fmt.Errorf("invalid card suit")
		}

	} else {
		return nil, fmt.Errorf("invalid card signature")
	}

	return c, nil
}

func (c *card) isSameSuit(o *card) bool {
	return c.suit == o.suit
}

func (c *card) isHigherThan(o *card) bool {
	return c.rank > o.rank
}

func (c *card) isEqual(o *card) bool {
	return c.rank == o.rank
}
