package poker

import (
	"fmt"
	"slices"
	"strings"
)

type hand struct {
	raw   string
	cards []*card
}

func NewHand(rawHand string) (*hand, error) {
	cards := strings.Split(rawHand, " ")
	if len(cards) != 5 {
		return nil, fmt.Errorf("invalid hand")
	}

	h := &hand{}
	h.raw = rawHand

	for i := range cards {
		c, err := NewCard(cards[i])
		if err != nil {
			return nil, err
		}

		h.cards = append(h.cards, c)
	}

	slices.SortFunc(h.cards, func(a *card, b *card) int {
		if a.isHigherThan(b) {
			return -1
		}

		if b.isHigherThan(a) {
			return 1
		}

		return 0
	})

	return h, nil
}

func (h *hand) getHighestRankCard(after *card) *card {
	var high *card
	for i := range h.cards {
		c := h.cards[i]
		if after != nil && (c.isEqual(after) || c.isHigherThan(after)) {
			continue
		}

		if high == nil || h.cards[i].isHigherThan(high) {
			high = h.cards[i]
			continue
		}
	}

	return high
}

func (h *hand) getHighestCombination(size int, after []*card) []*card {
	var comb []*card
	m := map[int][]*card{}

	for i := range h.cards {
		c := h.cards[i]
		if after != nil && c.isEqual(after[0]) {
			continue
		}

		if _, ok := m[c.rank]; !ok {
			m[c.rank] = []*card{}
		}

		m[c.rank] = append(m[c.rank], c)
	}

	highestPairValue := 0
	for i := range m {
		if len(m[i]) != size {
			continue
		}

		if m[i][0].rank <= highestPairValue {
			continue
		}

		highestPairValue = m[i][0].rank
		comb = []*card{}
		comb = append(comb, m[i]...)
	}

	return comb
}
