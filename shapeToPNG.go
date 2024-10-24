package main

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

// convert the grid shape into a slice of pixels
func PixelsFromShape(sh ShapeHolder, cellWidth, cellHeight int) [][]*imagehandling.Pixel {

	white := &imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}

	// make [][]*Pixel of all white pixels
	// Get number of Cells in row and number of rows
	// image will be number of Cells in row * cellWidth wide
	// image will be number of rows * cellHeight high
	g := sh.getShape()

	pixels := make([][]*imagehandling.Pixel, len(g.grid)*cellHeight)
	// loop
	for pixelY := 0; pixelY < len(pixels); pixelY++ {
		pixels[pixelY] = make([]*imagehandling.Pixel, len(g.grid[0])*cellWidth)

		for pixelX := 0; pixelX < len(pixels[pixelY]); pixelX++ {

			pixels[pixelY][pixelX] = white
		}
	}

	return pixels

}

func drawWalls(sh ShapeHolder, pixels [][]*imagehandling.Pixel, wallThickness int) [][]*imagehandling.Pixel {

	black := &imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	g := sh.getShape()

	cellWidth := len(pixels[0]) / g.columns
	cellHeight := len(pixels) / g.rows
	// worry about this later
	for row := 0; row < g.rows; row++ {
		for column := 0; column < g.columns; column++ {
			var c cell.Cell
			c2, ok := c.(*(cell.BaseCell))
			if !ok {
				return nil
			}

			if g.grid[row][column] == nil {
				// later fill in or ignore
				continue
			} else {
				c = g.grid[row][column]
				topLeft := []int{cellHeight * (row), cellWidth * (column)}
				bottomLeft := []int{cellHeight*(row+1) - 1, cellWidth * (column)}
				bottomRight := []int{cellHeight*(row+1) - 1, cellWidth*(column+1) - 1}
				topRight := []int{cellHeight * (row), cellWidth*(column+1) - 1}

				// clock-wise
				if !c.IsLinked(c2.West) {

					draw.StraightLine(bottomLeft[0], bottomLeft[1], topLeft[0], topLeft[1], pixels, *black)
				}

				if !c.IsLinked(c2.North) {

					draw.StraightLine(topRight[0], topRight[1], topLeft[0], topLeft[1], pixels, *black)
				}

				if !c.IsLinked(c2.East) {

					draw.StraightLine(bottomRight[0], bottomRight[1], topRight[0], topRight[1], pixels, *black)
				}

				if !c.IsLinked(c2.South) {

					draw.StraightLine(bottomRight[0], bottomRight[1], bottomLeft[0], bottomLeft[1], pixels, *black)
				}

			}

		}
	}

	return pixels
}
