package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

//	only algorithms like
//
// Wilson's, Recursive Backtracker, and Aldous-Broder will work with the library in its current state.

type TriangleGrid struct {
	*Shape
}

func CreateTriangleGrid(rows, columns int) *TriangleGrid {

	tg := &TriangleGrid{Shape: &Shape{
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}}

	prepareTriangleGrid(tg.getShape())
	configureTriangleCells(tg)

	return tg
}

func (tg *TriangleGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (tg *TriangleGrid) getShape() *Shape {
	return tg.Shape
}

func prepareTriangleGrid(g *Shape) {
	grid := make([][]cell.Cell, g.rows)

	for row := 0; row < g.rows; row++ {
		grid[row] = make([]cell.Cell, g.columns)

		for column := 0; column < g.columns; column++ {
			grid[row][column] = cell.CreateTriangleCell(row, column)
		}
	}

	g.grid = grid
}

func configureTriangleCells(tg *TriangleGrid) {

	grid := tg.getShape().grid

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {

			c := tg.GetCell(row, column).(*cell.TriangleCell)

			west := tg.GetCell(row, column-1)
			if west != nil {

				c.West = west.(*cell.TriangleCell)
			}

			east := tg.GetCell(row, column+1)
			if east != nil {

				c.East = east.(*cell.TriangleCell)
			}

			if c.Upright {

				south := tg.GetCell(row+1, column)
				if south != nil {
					c.South = south.(*cell.TriangleCell)
				}
			} else {

				north := tg.GetCell(row-1, column)
				if north != nil {
					c.North = north.(*cell.TriangleCell)
				}
			}

		}
	}
}

// this is going to need to be reworked
func (tg *TriangleGrid) toPNG(filepath string, size int) {
	shape := tg.getShape()

	halfWidth := float64(size) / 2.0
	height := float64(size) * math.Sqrt(3) / 2
	halfHeight := height / 2.0

	imgWidth := int(size * (shape.columns + 1) / 2.0)
	imgHeight := int(height) * shape.rows

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+1, imgHeight+1))

	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, white)
		}
	}

	for row := 0; row < len(tg.getShape().grid); row++ {
		for column := 0; column < len(tg.getShape().grid[row]); column++ {
			cell := tg.GetCell(row, column).(*cell.TriangleCell)

			cx := halfWidth + float64(column)*halfWidth
			cy := halfHeight + float64(row)*height

			westX := int(cx - halfWidth)
			midX := int(cx)
			eastX := int(cx + halfWidth)

			var apexY int
			var baseY int
			if cell.Upright {
				apexY = int(cy - halfHeight)
				baseY = int(cy + halfHeight)
			} else {
				apexY = int(cy + halfHeight)
				baseY = int(cy - halfHeight)
			}

			if !cell.IsLinked(cell.West) {
				draw.StraightLine2(westX, baseY, midX, apexY, img, black)
			}

			if !cell.IsLinked(cell.East) {
				draw.StraightLine2(eastX, baseY, midX, apexY, img, black)
			}

			noSouth := cell.Upright && !cell.IsLinked(cell.South)
			notLinked := !cell.Upright && !cell.IsLinked(cell.North)

			if row == 0 && column == 0 {

				fmt.Println(noSouth || notLinked)
			}
			if noSouth || notLinked {
				draw.StraightLine2(eastX, baseY, westX, baseY, img, black)
			}
		}

	}
	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)

}
