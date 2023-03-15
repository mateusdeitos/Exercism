type TAllergen = "eggs" |
	"peanuts" |
	"shellfish" |
	"strawberries" |
	"tomatoes" |
	"chocolate" |
	"pollen" |
	"cats";

const allergiesDict: Array<{ allergen: TAllergen, index: number }> = [
	{
		allergen: "cats",
		index: 128,
	},
	{
		allergen: "pollen",
		index: 64,
	},
	{
		allergen: "chocolate",
		index: 32,
	},
	{
		allergen: "tomatoes",
		index: 16,
	},
	{
		allergen: "strawberries",
		index: 8,
	},
	{
		allergen: "shellfish",
		index: 4,
	},
	{
		allergen: "peanuts",
		index: 2,
	},
	{
		allergen: "eggs",
		index: 1,
	},
];

const sumIndex = allergiesDict.reduce((sum, { index }) => sum + index, 0) + 1;

export class Allergies {
	private allergenIndex: number = 0;

	constructor(allergenIndex: number) {
		this.setAllergenIndex(allergenIndex);
	}

	private setAllergenIndex(allergenIndex: number): void {
		this.allergenIndex = allergenIndex;

		if (this.allergenIndex > sumIndex) {
			const timesHigher = Math.floor(this.allergenIndex / sumIndex);
			this.allergenIndex -= timesHigher * sumIndex;
		}
	}

	public list(): TAllergen[] {
		let score = this.allergenIndex;

		return allergiesDict.reduce<TAllergen[]>((result, { allergen, index }) => {
			if (score >= index) {
				score -= index;
				return [
					allergen,
					...result,
				];
			}

			return result;
		}, []);
	}

	public allergicTo(allergen: TAllergen): boolean {
		return this.list().includes(allergen);
	}
}
