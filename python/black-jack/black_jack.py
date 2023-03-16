from typing import Literal, NewType
"""Functions to help play and score a game of blackjack.

How to play blackjack:    https://bicyclecards.com/how-to-play/blackjack/
"Standard" playing cards: https://en.wikipedia.org/wiki/Standard_52-card_deck
"""

CardType = Literal["K", "Q", "J", "A", "2",
                   "3", "4", "5", "6", "7", "8", "9", "10"]


def value_of_card(
	card: CardType
) -> int:
	if (card in ["K", "Q", "J"]):
		return 10

	if (card == "A"):
		return 1

	return int(card)


def higher_card(card_one: CardType, card_two: CardType):
	value_one = value_of_card(card_one)
	value_two = value_of_card(card_two)
	if (value_one == value_two):
		return card_one, card_two

	if (value_one > value_two):
		return card_one

	return card_two


def value_of_ace(card_one: CardType, card_two: CardType):
	total_value = value_of_card(card_one) + value_of_card(card_two)
	has_ace = "A" in [card_one, card_two]
	blackjack = 21
	gap_to_blackjack = blackjack - total_value
	if (gap_to_blackjack > 10 and not has_ace):
		return 11

	return 1


def is_blackjack(card_one: CardType, card_two: CardType):
	if ("A" not in [card_one, card_two]):
		return False

	card_not_ace = card_one if card_one != "A" else card_two
	return value_of_card(card_not_ace) == 10


def can_split_pairs(card_one: CardType, card_two: CardType):
	return value_of_card(card_one) == value_of_card(card_two)


def can_double_down(card_one: CardType, card_two: CardType):
	total = value_of_card(card_one) + value_of_card(card_two)

	return total in [9, 10, 11]
