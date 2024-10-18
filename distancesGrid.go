package main

import (
	"strconv"
)

type DistancesGrid struct {
	*Distances
}

func CreateDistancesGrid(rows, columns int) *Grid {

	plainGrid := CreatePlainGrid(rows, columns)

	plainGrid.SpecialGrid = &DistancesGrid{
		Distances: nil,
	}

	return plainGrid
}

func (dg *DistancesGrid) cellContents(cell *Cell) string {
	if dg.Distances == nil {
		panic("distances were not initialized")
	} else {
		return strconv.Itoa(dg.Distances.cells[cell])
	}
}
