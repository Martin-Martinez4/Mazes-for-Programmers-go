package main

type ShapeHolder interface {
	getShape() *Shape
	toPNG(filepath string, cellSize int)
	ContentsOf(*Cell) string
}

func numberOfDeadEnds(sh ShapeHolder) int {
	grid := sh.getShape().grid
	rows := sh.getShape().rows
	columns := sh.getShape().columns

	deadEnds := 0

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if len(grid[row][column].links) == 1 {
				deadEnds++
			}
		}
	}

	return deadEnds
}
