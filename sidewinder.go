package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

func Sidewinder(sh ShapeHolder) {
	grid := sh.getShape()

	run := []cell.Cell{}

	var c cell.Cell

	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			c = grid.grid[row][column]
			c2, ok := c.(*cell.BaseCell)
			if !ok {
				return
			}
			run = append(run, c2)

			atEastBound := c2.East == nil
			atNorthBound := c2.North == nil

			shouldCloseOut := atEastBound || (!atNorthBound && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]
				m2, ok := c.(*cell.BaseCell)
				if !ok {
					return
				}

				if m2.North != nil {

					member.Link(m2.North)
				}

				run = []cell.Cell{}

			} else {
				c.Link(c2.East)
			}

		}
	}
}
