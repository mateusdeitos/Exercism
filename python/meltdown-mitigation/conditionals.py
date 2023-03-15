"""Functions to prevent a nuclear meltdown."""


from typing import Literal


def is_criticality_balanced(temperature: float, neutrons_emitted: float):
	temperature_check = temperature < 800
	neutrons_check = neutrons_emitted > 500
	product_check = (temperature * neutrons_emitted) < 500000

	if not temperature_check:
		return False

	if not neutrons_check:
		return False

	if not product_check:
		return False

	return True


def reactor_efficiency(
	voltage: int,
	current: int,
	theoretical_max_power: int
) -> Literal['green', 'orange', 'red', 'black']:
	generated_power = voltage * current
	efficiency = (generated_power / theoretical_max_power) * 100

	if efficiency >= 80:
		return 'green'

	if efficiency >= 60:
		return 'orange'

	if efficiency >= 30:
		return 'red'

	return 'black'


def fail_safe(
	temperature: int,
	neutrons_produced_per_second: int,
	threshold: int
) -> Literal["LOW", "NORMAL", "DANGER"]:
	criticality = temperature * neutrons_produced_per_second
	low_threshold = threshold * 0.9
	high_threshold = threshold * 1.1

	if criticality < low_threshold:
		return "LOW"

	if low_threshold <= criticality <= high_threshold:
		return "NORMAL"

	return "DANGER"


fail_safe(10, 901, 10000)
