package main

import (
	"image"
	"image/color"
	"math"

	drw "github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

//	only algorithms like
//
// Wilson's, Recursive Backtracker, and Aldous-Broder will work with the library in its current state.

type PolarGrid struct {
	Shape *Shape
}

func CreatePolarGrid(rows, columns int) *PolarGrid {
	shape := &Shape{
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}

	prepareGrid(shape)
	configureCells(shape)

	return &PolarGrid{Shape: shape}
}

func (pg *PolarGrid) ContentsOf(cell *Cell) string {
	return " "
}

func (pg *PolarGrid) getShape() *Shape {
	return pg.Shape
}

// this is going to need to be reworked
func (pg *PolarGrid) toPNG(filepath string, cellSize int) {
	shape := pg.getShape()
	grid := shape.grid

	imgSize := 2 * shape.rows * cellSize

	// background := imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}
	wall := imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgSize+1, imgSize+1))
	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imgSize+1; x++ {
		for y := 0; y < imgSize+1; y++ {
			img.Set(x, y, white)
		}
	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgSize+1, imgSize+1)

	center := imgSize / 2

	for row := 0; row < pg.getShape().rows; row++ {
		for column := 0; column < pg.getShape().columns; column++ {

			cell := grid[row][column]

			theta := 2 * math.Pi / float64(len(grid[cell.row]))
			innerRadius := float64(cell.row * cellSize)
			outerRadius := float64((cell.row + 1) * cellSize)
			thetaCcw := float64(cell.column) * theta
			thetaCw := float64(cell.column+1) * theta

			ax := center + int(innerRadius*math.Cos(thetaCcw))
			ay := center + int(innerRadius*math.Sin(thetaCcw))
			bx := center + int(outerRadius*math.Cos(thetaCcw))
			by := center + int(outerRadius*math.Sin(thetaCcw))

			cx := center + int(innerRadius*math.Cos(thetaCw))
			cy := center + int(innerRadius*math.Sin(thetaCw))
			dx := center + int(outerRadius*math.Cos(thetaCw))
			dy := center + int(outerRadius*math.Sin(thetaCw))

			if !cell.isLinked(cell.north) {
				drw.StraightLine(ax, ay, cx, cy, pixels, wall)
			}
			if !cell.isLinked(cell.east) {
				drw.StraightLine(cx, cy, dx, dy, pixels, wall)
			}
			if cell.row == shape.rows-1 {
				drw.StraightLine(bx, by, dx, dy, pixels, wall)
			}

		}
	}

	imagehandling.WritePNGFromPixels(filepath, pixels)

}
