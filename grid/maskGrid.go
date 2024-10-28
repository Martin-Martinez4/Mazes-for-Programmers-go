package grid

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

type MaskGrid struct {
	Shape *Shape
	Mask  *Mask
}

func CreateMaskGrid(mask *Mask) *MaskGrid {

	if mask == nil {

		panic("No mask provided")
	}
	rows := mask.rows
	columns := mask.columns

	shape := &Shape{
		Rows:    rows,
		Columns: columns,
		Size:    rows * columns,
	}

	prepareMaskGrid(shape, mask)
	configureCells(shape)

	return &MaskGrid{Shape: shape, Mask: mask}
}

func (pg *MaskGrid) ContentsOf(cell cell.Cell) string {
	return " "
}

func (pg *MaskGrid) GetShape() *Shape {
	return pg.Shape
}

func prepareMaskGrid(g *Shape, mask *Mask) {
	grid := make([][]cell.Cell, g.Rows)

	for row := 0; row < g.Rows; row++ {
		grid[row] = make([]cell.Cell, g.Columns)

		for column := 0; column < g.Columns; column++ {
			if mask.isOn(row, column) {

				grid[row][column] = cell.CreateBaseCell(row, column)
			} else {

				grid[row][column] = nil
			}
		}
	}

	g.Grid = grid
}

func (pg *MaskGrid) ToPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(pg, cellSize, cellSize)
	withWalls := drawWalls(pg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)
}
