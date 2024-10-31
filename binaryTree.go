package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func BinaryTree(sh grid.ShapeHolder) {

	grid := sh.GetShape()

	// replace this later
	for row := 0; row < grid.Rows(); row++ {

		for column := 0; column < grid.Columns(); column++ {
			c2, ok := (grid.Grid[row][column]).(*cell.BaseCell)
			if !ok {
				return
			}

			neighbors := []cell.Cell{}

			if c2.North != nil {
				neighbors = append(neighbors, c2.North)
			}

			if c2.East != nil {
				neighbors = append(neighbors, c2.East)
			}

			if len(neighbors) == 0 {
				continue
			} else {
				randIndex := rand.Intn(len(neighbors))
				c2.Link(neighbors[randIndex])
			}

		}
	}
}
