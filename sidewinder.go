package main

import (
	"math/rand"
)

func Sidewinder(sh ShapeHolder) {
	grid := sh.getShape()

	run := []Cell{}

	var cell Cell

	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			cell = grid.grid[row][column]
			c2, ok := cell.(*BaseCell)
			if !ok {
				return
			}
			run = append(run, c2)

			atEastBound := c2.east == nil
			atNorthBound := c2.north == nil

			shouldCloseOut := atEastBound || (!atNorthBound && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]
				m2, ok := cell.(*BaseCell)
				if !ok {
					return
				}

				if m2.north != nil {

					member.Link(m2.north)
				}

				run = []Cell{}

			} else {
				cell.Link(c2.east)
			}

		}
	}
}
