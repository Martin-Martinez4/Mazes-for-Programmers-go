package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func GrowingTree(sh grid.ShapeHolder, getCell func(cells []cell.Cell) cell.Cell) {
	grid := sh.GetShape().Grid

	randRow := rand.Intn(len(grid))
	randCol := rand.Intn(len(grid[randRow]))

	active := []cell.Cell{grid[randRow][randCol]}

	for len(active) > 0 {
		c := getCell(active)

		neighbor := randomUnlinked(c.Neighbors())

		if neighbor != nil {
			c.Link(neighbor)
			active = append(active, neighbor)
		} else {

			// active = append(active[:randIndex], active[randIndex+1:]...)
			active = remove(c, active)
		}
	}
}

func remove(c cell.Cell, cells []cell.Cell) []cell.Cell {
	for i := 0; i < len(cells); i++ {
		if cells[i] == c {
			cells = append(cells[:i], cells[i+1:]...)

			return cells
		}
	}

	return cells
}
