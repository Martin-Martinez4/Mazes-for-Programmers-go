package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

// Randomw walk like in aldousBroder
// but when a loop happens look for a unvisited cell with a visited neighbor

func HuntAndKill(sh grid.ShapeHolder) {
	current := sh.GetShape().RandomCell()
	// grid := sh.GetShape().Grid
	// rows := sh.GetShape().Rows
	// columns := sh.GetShape().Columns

	for current != nil {
		unvisitedNeighbors := unvisitedNeighbors(current)

		if len(unvisitedNeighbors) != 0 {
			neighbor := unvisitedNeighbors[rand.Intn(len(unvisitedNeighbors))]
			current.Link(neighbor)
			current = neighbor
		} else {
			current = nil

			// for row := 0; row < rows; row++ {
			// 	for column := 0; column < columns; column++ {

			for c := range sh.EachCell() {

				// row := c.Row()
				// column := c.Column()

				if c != nil {

					visitedNeighbors := visitedNeighbors(c)

					if (len(c.Links()) == 0) && (len(visitedNeighbors) > 0) {
						current = c

						neighbor := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
						current.Link(neighbor)

						break
					}
				}
			}

			// 	}
			// }
		}

	}

}

func unvisitedNeighbors(c cell.Cell) []cell.Cell {

	var unvisited []cell.Cell
	neighbors := c.Neighbors()

	for i := 0; i < len(neighbors); i++ {
		if len(neighbors[i].Links()) == 0 {
			unvisited = append(unvisited, neighbors[i])
		}
	}

	return unvisited
}

func visitedNeighbors(c cell.Cell) []cell.Cell {

	var visited []cell.Cell
	neighbors := c.Neighbors()

	for i := 0; i < len(neighbors); i++ {
		if len(neighbors[i].Links()) > 0 {
			visited = append(visited, neighbors[i])
		}
	}

	return visited
}
