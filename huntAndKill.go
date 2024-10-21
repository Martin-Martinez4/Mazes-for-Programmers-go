package main

import "math/rand"

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
			current.link(neighbor)
			current = neighbor
		} else {
			current = nil

			for row := 0; row < rows; row++ {
				for column := 0; column < columns; column++ {

					cell := grid[row][column]

					visitedNeighbors := visitedNeighbors(cell)

					if (len(cell.links) == 0) && (len(visitedNeighbors) > 0) {
						current = cell

						neighbor := visitedNeighbors[rand.Intn(len(visitedNeighbors))]
						current.link(neighbor)

						break
					}
				}
			}
		}

	}

}

func unvisitedNeighbors(cell *Cell) []*Cell {

	var unvisited []*Cell
	neighbors := cell.neighbors()

	for i := 0; i < len(neighbors); i++ {
		if len(neighbors[i].links) == 0 {
			unvisited = append(unvisited, neighbors[i])
		}
	}

	return unvisited
}

func visitedNeighbors(cell *Cell) []*Cell {

	var visited []*Cell
	neighbors := cell.neighbors()

	for i := 0; i < len(neighbors); i++ {
		if len(neighbors[i].links) > 0 {
			visited = append(visited, neighbors[i])
		}
	}

	return visited
}
