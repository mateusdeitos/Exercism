const dict = [
	{ binary: 1000, value: 'jump', },
	{ binary: 100, value: 'close your eyes', },
	{ binary: 10, value: 'double blink', },
	{ binary: 1, value: 'wink', },
]

const reverseValue = 10000;

export function commands(decimal: number) {
	let binary = decimalToBinary(decimal);
	if (binary <= 0) {
		return [];
	}

	const reverse = binary - reverseValue >= 0;
	if (reverse) {
		binary -= reverseValue;
	}

	const operations: string[] = [];
	let stop: boolean;
	do {
		stop = true;
		dict.forEach(({ binary: n, value }) => {
			if (binary - n >= 0) {
				stop = false;
				binary -= n;
				operations.unshift(value);
			}
		});

	} while (binary > 0 && !stop);

	if (reverse) {
		return operations.reverse();
	}

	return operations;
}

const decimalToBinary = (decimal: number): number => {
	return Number((decimal >>> 0).toString(2));
}
