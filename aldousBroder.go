package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

// Simple random walk
func AldousBroder(sh grid.ShapeHolder) {
	cell := sh.GetShape().RandomCell()
	unvisited := sh.GetShape().Size - 1

	for unvisited > 0 {
		neighbors := cell.Neighbors()
		randIndex := rand.Intn(len(neighbors))

		neighbor := neighbors[randIndex]

		if len(neighbor.Links()) == 0 {
			cell.Link(neighbor)
			unvisited--
		}

		cell = neighbor
	}

}
