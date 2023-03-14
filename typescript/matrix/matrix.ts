export class Matrix {
	private matrix: number[][] = [];
	constructor(matrix: string) {
		matrix.split("\n").forEach((row, rowIndex) => {
			this.matrix[rowIndex] = row.split(" ").map(Number);
		});
	}

	get rows(): number[][] {
		return this.matrix;
	}

	get columns(): number[][] {
		const columns: number[][] = [];
		this.matrix.forEach((row, rowIndex) => {
			row.forEach((column, columnIndex) => {
				if (!columns[columnIndex]) {
					columns[columnIndex] = [];
				}
				columns[columnIndex][rowIndex] = column;
			});
		});
		return columns;
	}

}
