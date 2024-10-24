package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"

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

func (pg *PlainGrid) ContentsOf(cell Cell) string {
	return " "
}

func (pg *PlainGrid) getShape() *Shape {
	return pg.Shape
}

func (pg *PlainGrid) toPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(pg, cellSize, cellSize)
	withWalls := drawWalls(pg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)

}
