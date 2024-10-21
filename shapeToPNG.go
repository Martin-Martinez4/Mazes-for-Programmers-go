package main

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

// convert the grid shape into a slice of pixels
func PixelsFromShape(sh ShapeHolder, cellWidth, cellHeight int) [][]*imagehandling.Pixel {

	white := &imagehandling.Pixel{255, 255, 255, 255}

	// make [][]*Pixel of all white pixels
	// Get number of Cells in row and number of rows
	// image will be number of Cells in row * cellWidth wide
	// image will be number of rows * cellHeight high
	g := sh.getShape()

	pixels := make([][]*imagehandling.Pixel, len(g.grid)*cellHeight)
	// loop
	for pixelY := 0; pixelY < len(pixels); pixelY++ {
		pixels[pixelY] = make([]*imagehandling.Pixel, len(g.grid[0])*cellWidth)

		for pixelX := 0; pixelX < len(pixels[0]); pixelX++ {
			pixels[pixelY][pixelX] = white
		}
	}

	return pixels

	// worry about this later
	// for row := 0; row < g.rows; row++ {
	// 	for column := 0; column < g.columns; column++ {
	// 		var cell *Cell

	// 		if g.grid[row][column] == nil {
	// 			// later fill in or ignore
	// 			cell = nil
	// 		} else {
	// 			cell = g.grid[row][column]
	// 		}

	// 		topLeft := []int{cellWidth * (column), cellHeight * (row)}
	// 		bottomLeft := []int{cellWidth * (column), cellHeight * (row + 1)}
	// 		bottomRight := []int{cellWidth * (column + 1), cellHeight * (row + 1)}
	// 		topRight := []int{cellWidth * (column + 1), cellHeight * (row)}

	// 		fmt.Println("top left: ", topLeft[0], topLeft[1])
	// 		fmt.Println("bottom left: ", bottomLeft[0], bottomLeft[1])
	// 		fmt.Println("bottom right: ", bottomRight[0], bottomRight[1])
	// 		fmt.Println("top right: ", topRight[0], topRight[1])

	// 		if cell.west == nil {
	// 			// draw west wall
	// 		} else {

	// 			ok := cell.links[cell.west]
	// 			if !ok {
	// 				// no link means wall
	// 				// draw west wall
	// 			}
	// 		}
	// 		if cell.north == nil {
	// 			// draw north wall
	// 		} else {

	// 			ok := cell.links[cell.north]
	// 			if !ok {
	// 				// no link means wall
	// 				// draw north wall
	// 			}
	// 		}

	// 		if cell.east == nil {
	// 			// draw east wall
	// 		} else {

	// 			ok := cell.links[cell.east]
	// 			if !ok {
	// 				// no link means wall
	// 				// draw east wall
	// 			}
	// 		}

	// 		if cell.south == nil {
	// 			// draw south wall
	// 		} else {

	// 			ok := cell.links[cell.south]
	// 			if !ok {
	// 				// no link means wall
	// 				// draw south wall
	// 			}
	// 		}

	// 	}
	// }

	return nil
}
