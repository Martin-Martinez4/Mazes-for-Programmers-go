package main

type DistancesGrid struct {
	*Distances
	Shape *Shape
}

func CreateDistancesGrid(rows, columns int) *DistancesGrid {

	shape := CreateShape(rows, columns)

	return &DistancesGrid{Shape: shape}
}

func (dg *DistancesGrid) setDistancesTo(row, column int) {
	dg.Distances = dg.Shape.grid[row][column].distances()
}

func (dg *DistancesGrid) ContentsOf(cell *Cell) string {
	if dg.Distances == nil {
		panic("distances were not initialized")
	} else {
		if dg.Distances.cells[cell] == STARTPOINT {
			return "@"
		} else {

			return string(IntToBase62(dg.Distances.cells[cell]))
		}
	}
}

func (dg *DistancesGrid) getShape() *Shape {
	return dg.Shape
}
