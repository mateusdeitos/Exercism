package poker

import (
	"fmt"
)

type rankingCategory int

const (
	HighCard      rankingCategory = 1 << iota
	OnePair       rankingCategory = 1 << iota
	TwoPair       rankingCategory = 1 << iota
	ThreeOfAKind  rankingCategory = 1 << iota
	Straight      rankingCategory = 1 << iota
	Flush         rankingCategory = 1 << iota
	FullHouse     rankingCategory = 1 << iota
	FourOfAKind   rankingCategory = 1 << iota
	StraightFlush rankingCategory = 1 << iota
)

func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return nil, fmt.Errorf("no hands informed")
	}

	scores := map[rankingCategory][]*hand{}

	for i := range hands {
		h, err := NewHand(hands[i])
		if err != nil {
			return nil, err
		}

		e := evaluator{h}
		r := e.getRankingOfHand()
		if _, ok := scores[r]; !ok {
			scores[r] = []*hand{}
		}

		scores[r] = append(scores[r], h)
	}

	var highestScore rankingCategory
	for score := range scores {
		if score > highestScore {
			highestScore = score
		}
	}

	return untieHands(highestScore, scores[highestScore]), nil
}
