package main

import (
	"math/rand"
)

func BinaryTree(grid *Grid) {
	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			cell := grid.grid[row][column]
			neighbors := []*Cell{}

			if cell.north != nil {
				neighbors = append(neighbors, cell.north)
			}

			if cell.east != nil {
				neighbors = append(neighbors, cell.east)
			}

			if len(neighbors) == 0 {
				continue
			} else {
				randIndex := rand.Intn(len(neighbors))
				cell.link(neighbors[randIndex])
			}

		}
	}
}
