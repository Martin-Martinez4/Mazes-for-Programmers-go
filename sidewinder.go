package main

import (
	"math/rand"
)

func Sidewinder(sh ShapeHolder) {
	grid := sh.getShape()

	run := []*Cell{}

	var cell *Cell

	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			cell = grid.grid[row][column]
			run = append(run, cell)

			atEastBound := cell.east == nil
			atNorthBound := cell.north == nil

			shouldCloseOut := atEastBound || (!atNorthBound && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]

				if member.north != nil {

					member.link(member.north)
				}

				run = []*Cell{}

			} else {
				cell.link(cell.east)
			}

		}
	}
}
