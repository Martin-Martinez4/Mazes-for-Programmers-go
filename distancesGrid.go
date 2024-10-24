package main

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

const STARTPOINT = 64

type DistancesGrid struct {
	*cell.Distances
	Shape *Shape
}

func CreateDistancesGrid(rows, columns int) *DistancesGrid {

	shape := CreateShape(rows, columns)

	return &DistancesGrid{Shape: shape}
}

func (dg *DistancesGrid) setDistancesTo(row, column int) {
	dg.Distances = dg.Shape.grid[row][column].Distances()
}

func (dg *DistancesGrid) ContentsOf(cell cell.Cell) string {
	if dg.Distances == nil {
		panic("distances were not initialized")
	} else {
		if dg.Distances.Cells[cell] == STARTPOINT {
			return "@"
		} else {

			return string(IntToBase62(dg.Distances.Cells[cell]))
		}
	}
}

func (dg *DistancesGrid) getShape() *Shape {
	return dg.Shape
}

func (dg *DistancesGrid) toPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(dg, cellSize, cellSize)
	withWalls := drawWalls(dg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)
}
