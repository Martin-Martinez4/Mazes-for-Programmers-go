package main

import (
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

	black := &imagehandling.Pixel{R: 20, G: 0, B: 0, A: 255}

	g := sh.getShape()

	cellWidth := len(pixels[0]) / g.columns
	cellHeight := len(pixels) / g.rows
	// worry about this later
	for row := 0; row < g.rows; row++ {
		for column := 0; column < g.columns; column++ {
			var cell *Cell

			if g.grid[row][column] == nil {
				// later fill in or ignore
				continue
			} else {
				cell = g.grid[row][column]
			}

			topLeft := []int{cellHeight * (row), cellWidth * (column)}
			bottomLeft := []int{cellHeight*(row+1) - 1, cellWidth * (column)}
			// bottomRight := []int{cellWidth * (column + 1), cellHeight * (row + 1)}
			topRight := []int{cellHeight * (row), cellWidth*(column+1) - 1}

			// fmt.Println("top left: ", topLeft[0], topLeft[1])
			// fmt.Println("bottom left: ", bottomLeft[0], bottomLeft[1])
			// fmt.Println("bottom right: ", bottomRight[0], bottomRight[1])
			// fmt.Println("top right: ", topRight[0], topRight[1])

			if cell.west == nil {
				// draw west wall
				for h := 0; h < cellHeight; h++ {
					for b := 0; b < wallThickness; b++ {
						pixels[topLeft[0]+h][topLeft[1]+b] = black
					}
				}
			} else {

				ok := cell.links[cell.west]
				if !ok {
					// no link means wall
					// draw west wall
					for h := 0; h < cellHeight; h++ {
						for b := 0; b < wallThickness; b++ {
							pixels[topLeft[0]+h][topLeft[1]+b] = black
						}
					}
				}
			}

			if cell.north == nil {
				// draw north wall
				for w := 0; w < cellWidth; w++ {
					for b := 0; b < wallThickness; b++ {
						pixels[topLeft[0]+b][topLeft[1]+w] = black
					}
				}
			} else {

				ok := cell.links[cell.north]
				if !ok {
					// no link means wall
					// draw north wall
					for w := 0; w < cellWidth; w++ {
						for b := 0; b < wallThickness; b++ {
							pixels[topLeft[0]+b][topLeft[1]+w] = black
						}
					}
				}
			}

			if cell.east == nil {
				// draw east wall
				for h := 0; h < cellHeight; h++ {
					for b := 0; b < wallThickness; b++ {
						pixels[topRight[0]+h][topRight[1]-b] = black
					}
				}
			} else {

				ok := cell.links[cell.east]
				if !ok {
					// no link means wall
					// draw east wall
					for h := 0; h < cellHeight; h++ {
						for b := 0; b < wallThickness; b++ {
							pixels[topRight[0]+h][topRight[1]-b] = black
						}
					}
				}
			}

			if cell.south == nil {
				// draw south wall
				for w := 0; w < cellWidth; w++ {
					for b := 0; b < wallThickness; b++ {
						pixels[bottomLeft[0]-b][bottomLeft[1]+w] = black
					}
				}
			} else {

				ok := cell.links[cell.south]
				if !ok {
					// no link means wall
					// draw south wall
					for w := 0; w < cellWidth; w++ {
						for b := 0; b < wallThickness; b++ {
							pixels[bottomLeft[0]+b][bottomLeft[1]+w] = black
						}
					}
				}
			}

		}
	}

	return pixels
}
