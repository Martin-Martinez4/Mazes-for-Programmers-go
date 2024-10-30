package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func SimplifiedPrims(sh grid.ShapeHolder) {
	grid := sh.GetShape().Grid

	randRow := rand.Intn(len(grid))
	randCol := rand.Intn(len(grid[randRow]))

	active := []cell.Cell{grid[randRow][randCol]}

	for len(active) > 0 {
		randIndex := rand.Intn(len(active))
		c := active[randIndex]

		neighbor := randomUnlinked(c.Neighbors())

		if neighbor != nil {
			c.Link(neighbor)
			active = append(active, neighbor)
		} else {

			active = append(active[:randIndex], active[randIndex+1:]...)
		}
	}

}

func Prims(sh grid.ShapeHolder) {
	grid := sh.GetShape().Grid

	// Give each cell a random weight
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			grid[row][column].SetWeight(rand.Intn(100))
		}
	}

	randRow := rand.Intn(len(grid))
	randCol := rand.Intn(len(grid[randRow]))

	active := cell.PriorityQueue{}
	active.Push(grid[randRow][randCol])

	for len(active) > 0 {
		c := active.Peek()

		neighbor := randomUnlinked(c.Neighbors())

		if neighbor != nil {
			c.Link(neighbor)
			active = append(active, neighbor)
		} else {

			active.Pop()
		}
	}

}

func randomUnlinked(cells []cell.Cell) cell.Cell {
	grid.Shuffle(cells)

	for i := 0; i < len(cells); i++ {
		if len(cells[i].Links()) == 0 {
			return cells[i]
		}
	}

	return nil
}
