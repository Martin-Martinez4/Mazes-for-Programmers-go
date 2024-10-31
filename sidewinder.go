package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func Sidewinder(sh grid.ShapeHolder) {
	grid := sh.GetShape()

	run := []cell.Cell{}

	var c cell.Cell

	// replace this later
	for row := 0; row < sh.Rows(); row++ {

		for column := 0; column < sh.Columns(); column++ {
			c = grid.Grid[row][column]
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
