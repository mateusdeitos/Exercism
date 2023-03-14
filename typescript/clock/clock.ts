const HOURS_IN_DAY = 24;
const MINUTES_IN_HOUR = 60;

export class Clock {
	private hours: number = 0;
	private minutes: number = 0;
	constructor(hour: number, minute: number = 0) {
		const extraHours = this.setMinutes(minute);
		this.setHours(hour + extraHours);
	}

	public toString(): string {
		return `${this.hours.toString().padStart(2, '0')}:${this.minutes.toString().padStart(2, '0')}`;
	}

	public plus(minutes: number): Clock {
		const extraHours = this.setMinutes(this.minutes + minutes);
		this.setHours(this.hours + extraHours);
		return this;
	}

	public minus(minutes: number): Clock {
		const extraHours = this.setMinutes(this.minutes - minutes);
		this.setHours(this.hours + extraHours);
		return this;
	}

	public equals(other: Clock): boolean {
		return this.toString() === other.toString();
	}

	private setHours(hours: number): void {
		const adjusted = (hours % HOURS_IN_DAY);
		this.hours = adjusted < 0 ? adjusted + HOURS_IN_DAY : adjusted;
	}

	/**
	 * Parse the minutes and return the amount of hours that should be added to the hours
	 */
	private setMinutes(minutes: number): number {
		if (minutes >= 0 && minutes < MINUTES_IN_HOUR) {
			this.minutes = minutes;
			return 0;
		}

		const adjusted = (minutes % MINUTES_IN_HOUR);
		const extraHours = Math.floor(minutes / MINUTES_IN_HOUR);

		this.minutes = adjusted < 0 ? adjusted + MINUTES_IN_HOUR : adjusted;
		return extraHours;
	}
}
