package poker

type suit string

const (
	diamond suit = "♢"
	spades  suit = "♤"
	hearts  suit = "♡"
	clubs   suit = "♧"
)

func (s *suit) isValid() bool {
	return *s == diamond ||
		*s == spades ||
		*s == hearts ||
		*s == clubs
}

const ACE_UPPER_BOUND = 14
const ACE_LOWER_BOUND = 1
