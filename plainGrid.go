package main

type PlainGrid struct {
	Shape *Shape
}

func CreatePlainGrid(rows, columns int) *PlainGrid {
	shape := &Shape{
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}

	prepareGrid(shape)
	configureCells(shape)

	return &PlainGrid{Shape: shape}
}

func (pg *PlainGrid) ContentsOf(cell *Cell) string {
	return " "
}

func (pg *PlainGrid) getShape() *Shape {
	return pg.Shape
}
