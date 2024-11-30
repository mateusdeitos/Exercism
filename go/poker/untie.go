package poker

// untieHands returns a slice of strings representing the winning hands based on the given ranking category and a slice of hands.
//
// Parameters:
// - rc: a rankingCategory representing the category of the hands.
// - hands: a slice of pointers to hand objects representing the hands to be untied.
//
// Return:
// - score: a slice of strings representing the winning hands.
func untieHands(rc rankingCategory, hands []*hand) []string {
	score := []string{}
	if len(hands) == 1 {
		score = append(score, hands[0].raw)
		return score
	}

	winners := []*hand{}

	switch rc {
	case Flush, StraightFlush, Straight:
		winners = filterHandsWithHighestCard(hands)
	case FullHouse:
		winners = filterHandsWithHighestCombinationOf(3, hands)

		if len(winners) > 1 {
			winners = filterHandsWithHighestCombinationOf(2, winners)
		}
	case FourOfAKind:
		winners = filterHandsWithHighestCombinationOf(4, hands)

		if len(winners) > 1 {
			winners = filterHandsWithHighestCard(winners)
		}
	case OnePair, TwoPair:
		winners = filterHandsWithHighestCombinationOf(2, hands)

		if len(winners) > 1 {
			winners = filterHandsWithHighestCard(winners)
		}
	case ThreeOfAKind:
		winners = filterHandsWithHighestCombinationOf(3, hands)

		if len(winners) > 1 {
			winners = filterHandsWithHighestCard(winners)
		}
	case HighCard:
		winners = filterHandsWithHighestCard(hands)
	}

	for i := range winners {
		h := winners[i]
		score = append(score, h.raw)
	}

	return score
}

// filterHandsWithHighestCombinationOf compares the highest combination of cards in a slice of hands and returns a slice of hands with the highest ranking cards.
//
// Parameters:
// - cardCombinationCount: an integer representing the number of cards in each combination.
// - hands: a slice of pointers to hand objects.
//
// Return:
// - a slice of pointers to hand objects with the highest ranking cards.
func filterHandsWithHighestCombinationOf(cardCombinationCount int, hands []*hand) []*hand {
	var highestSet []*card
	winners := []*hand{}

	rounds := int(len(hands) / cardCombinationCount)
	if rounds == 0 {
		rounds = 1
	}

	for i := 0; i < rounds; i++ {
		var roundHighestPair []*card
		for i := range hands {
			h := hands[i]
			high := h.getHighestCombination(cardCombinationCount, highestSet)
			if high == nil {
				winners = append(winners, h)
				continue
			}

			if roundHighestPair == nil || high[0].isHigherThan(roundHighestPair[0]) {
				roundHighestPair = high
				winners = []*hand{h}
				continue
			}

			if high[0].isEqual(roundHighestPair[0]) {
				winners = append(winners, h)
			}
		}

		if len(winners) == 1 {
			break
		}

		highestSet = roundHighestPair
		hands = winners
	}

	return winners
}

func filterHandsWithHighestCard(hands []*hand) []*hand {
	var highestRankingCard *card
	handsWithHighestRankingCard := []*hand{}

	for i := 0; i < 5; i++ {
		var roundHighestRankingCard *card
		for i := range hands {
			h := hands[i]
			high := h.getHighestRankCard(highestRankingCard)
			if high == nil {
				handsWithHighestRankingCard = append(handsWithHighestRankingCard, h)
				continue
			}

			if roundHighestRankingCard == nil || high.isHigherThan(roundHighestRankingCard) {
				roundHighestRankingCard = high
				handsWithHighestRankingCard = []*hand{h}
				continue
			}

			if high.isEqual(roundHighestRankingCard) {
				handsWithHighestRankingCard = append(handsWithHighestRankingCard, h)
			}
		}

		if len(handsWithHighestRankingCard) == 1 {
			break
		}

		highestRankingCard = roundHighestRankingCard
		hands = handsWithHighestRankingCard
	}

	return handsWithHighestRankingCard
}
