def exchange_money(budget: float, exchange_rate: float):
	return budget / exchange_rate


def get_change(budget: float, exchanging_value: float):
	return budget - exchanging_value


def get_value_of_bills(denomination: int, number_of_bills: int):
	return denomination * number_of_bills


def get_number_of_bills(budget: float, denomination: int):
	return budget // denomination


def get_leftover_of_bills(budget: float, denomination: int):
	return budget % denomination


def exchangeable_value(budget: float, exchange_rate: float, spread: int, denomination: int):
	exchange_rate_with_fee = exchange_rate * (1 + spread / 100)
	exchanging_value = exchange_money(budget, exchange_rate_with_fee)
	number_of_bills = get_number_of_bills(exchanging_value, denomination)
	return get_value_of_bills(denomination, number_of_bills)
