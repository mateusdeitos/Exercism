package poker

import (
	"math"
	"sort"
)

type evaluator struct {
	hand *hand
}

func (e *evaluator) getRankingOfHand() rankingCategory {
	isFlush := e.isFlush()
	isStraight := e.isStraight()

	if isFlush && isStraight {
		return StraightFlush
	}

	if !isStraight && e.isFourOfAKind() {
		return FourOfAKind
	}

	if !isStraight && e.isFullHouse() {
		return FullHouse
	}

	if isFlush {
		return Flush
	}

	if isStraight {
		return Straight
	}

	if e.isThreeOfAKind() {
		return ThreeOfAKind
	}

	if e.isTwoPair() {
		return TwoPair
	}

	if e.isOnePair() {
		return OnePair
	}

	return HighCard
}

func (e *evaluator) isOnePair() bool {
	m := map[int]int{}

	for _, c := range e.hand.cards {
		m[c.rank]++
	}

	pairCount := 0

	for _, v := range m {
		if v == 2 {
			pairCount++
		}
	}

	return pairCount == 1
}

func (e *evaluator) isTwoPair() bool {
	m := map[int]int{}

	for _, c := range e.hand.cards {
		m[c.rank]++
	}

	pairCount := 0

	for _, v := range m {
		if v == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func (e *evaluator) isThreeOfAKind() bool {
	m := map[int]int{}

	for _, c := range e.hand.cards {
		m[c.rank]++
	}

	for _, v := range m {
		if v == 3 {
			return true
		}
	}

	return false
}

func (e *evaluator) isFullHouse() bool {
	m := map[int]int{}

	for _, c := range e.hand.cards {
		if _, ok := m[c.rank]; !ok {
			m[c.rank] = 0

			if len(m) > 2 {
				return false
			}
		}

		m[c.rank]++
	}

	for _, v := range m {
		if v != 2 && v != 3 {
			return false
		}
	}

	return true
}

func (e *evaluator) isFourOfAKind() bool {
	m := map[int]int{}

	for _, c := range e.hand.cards {
		m[c.rank]++
	}

	for _, v := range m {
		if v == 4 {
			return true
		}
	}

	return false
}

func (e *evaluator) isFlush() bool {
	for _, c := range e.hand.cards {
		if !c.isSameSuit(e.hand.cards[0]) {
			return false
		}
	}

	return true
}

func (e *evaluator) isStraight() bool {

	fn := func(cards []*card) (bool, bool) {
		var prev *card
		invalid := false
		hasAce := false

		for i := range cards {
			c := cards[i]
			if !hasAce && c.rank == ACE_UPPER_BOUND {
				hasAce = true
			}

			if i == 0 {
				prev = c
				continue
			}

			diff := int(math.Abs(float64(prev.rank - c.rank)))
			if diff != 1 {
				invalid = true
				break
			}

			prev = c
		}

		return invalid, hasAce
	}

	invalid, hasAce := fn(e.hand.cards)
	if !hasAce || !invalid {
		return !invalid
	}

	// do the same comparison but sort it considering that Ace is 1
	sorted := make([]*card, len(e.hand.cards))

	copy(sorted, e.hand.cards)

	ace := *sorted[0]
	ace.rank = ACE_LOWER_BOUND
	sorted[0] = &ace

	sort.Slice(sorted, func(i, j int) bool {
		iRank := sorted[i].rank
		jRank := sorted[j].rank

		if iRank == ACE_UPPER_BOUND {
			iRank = ACE_LOWER_BOUND
		}

		if jRank == ACE_UPPER_BOUND {
			jRank = ACE_LOWER_BOUND
		}

		return iRank > jRank
	})

	invalid, _ = fn(sorted)

	if !invalid {
		e.hand.cards = sorted
	}

	return !invalid
}
