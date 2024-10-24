package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

// Randomw walk like in aldousBroder
// but when a loop happens look for a unvisited cell with a visited neighbor

func HuntAndKill(sh ShapeHolder) {
	current := sh.getShape().randomCell()
	grid := sh.getShape().grid
	rows := sh.getShape().rows
	columns := sh.getShape().columns

	for current != nil {
		unvisitedNeighbors := unvisitedNeighbors(current)

		if len(unvisitedNeighbors) != 0 {
			neighbor := unvisitedNeighbors[rand.Intn(len(unvisitedNeighbors))]
			current.Link(neighbor)
			current = neighbor
		} else {
			current = nil

			for row := 0; row < rows; row++ {
				for column := 0; column < columns; column++ {

					cell := grid[row][column]

					if cell != nil {

						visitedNeighbors := visitedNeighbors(cell)

						if (len(cell.Links()) == 0) && (len(visitedNeighbors) > 0) {
							current = cell

							neighbor := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
							current.Link(neighbor)

							break
						}
					}

				}
			}
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
