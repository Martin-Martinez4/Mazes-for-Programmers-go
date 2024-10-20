package main

import "math/rand"

// Simple random walk
func AldousBroder(sh ShapeHolder) {
	cell := sh.getShape().randomCell()
	unvisited := sh.getShape().size - 1

	for unvisited > 0 {
		neighbors := cell.neighbors()
		randIndex := rand.Intn(len(neighbors))

		neighbor := neighbors[randIndex]

		if len(neighbor.links) == 0 {
			cell.link(neighbor)
			unvisited--
		}

		cell = neighbor
	}

}
