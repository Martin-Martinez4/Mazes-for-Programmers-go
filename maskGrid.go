package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"

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
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}

	prepareMaskGrid(shape, mask)
	configureCells(shape)

	return &MaskGrid{Shape: shape, Mask: mask}
}

func (pg *MaskGrid) ContentsOf(cell *Cell) string {
	return " "
}

func (pg *MaskGrid) getShape() *Shape {
	return pg.Shape
}

func prepareMaskGrid(g *Shape, mask *Mask) {
	grid := make([][]*Cell, g.rows)

	for row := 0; row < g.rows; row++ {
		grid[row] = make([]*Cell, g.columns)

		for column := 0; column < g.columns; column++ {
			if mask.isOn(row, column) {

				grid[row][column] = CreateCell(row, column)
			} else {

				grid[row][column] = nil
			}
		}
	}

	g.grid = grid
}

func (pg *MaskGrid) toPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(pg, cellSize, cellSize)
	withWalls := drawWalls(pg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)
}
