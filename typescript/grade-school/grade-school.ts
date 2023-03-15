const clone = <T>(obj: T) => JSON.parse(JSON.stringify(obj)) as T;

export class GradeSchool {
	private readonly students: string[] = [];
	private readonly _roster: { [key: number]: string[] } = {}
	roster() {
		return clone(this._roster);
	}

	add(name: string, grade: number) {
		if (this.students.includes(name)) {
			return this.removeGradeOfStudent(name);
		}

		if (!this._roster[grade]) {
			this._roster[grade] = []
		}

		this.students.push(name);
		this._roster[grade].push(name);
		this._roster[grade].sort();
	}

	grade(grade: number) {
		if (!this._roster[grade]) {
			return [];
		}

		return clone(this._roster[grade]).sort();
	}

	private removeGradeOfStudent(name: string) {
		for (const grade in this._roster) {
			if (!this._roster[grade].includes(name)) {
				continue;
			}

			this._roster[grade] = this._roster[grade].filter(student => student !== name);
		}
	}
}
