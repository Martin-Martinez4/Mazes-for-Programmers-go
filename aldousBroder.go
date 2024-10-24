package main

import "math/rand"

// Simple random walk
func AldousBroder(sh ShapeHolder) {
	cell := sh.getShape().randomCell()
	unvisited := sh.getShape().size - 1

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
