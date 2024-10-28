package grid

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

type ShapeHolder interface {
	GetShape() *Shape
	ToPNG(filepath string, cellSize int)
	ContentsOf(cell.Cell) string
}

func deadEnds(sh ShapeHolder) []cell.Cell {
	grid := sh.GetShape().Grid
	rows := sh.GetShape().Rows
	columns := sh.GetShape().Columns

	deadEnds := []cell.Cell{}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if grid[row][column] == nil {
				continue
			}
			if len(grid[row][column].Links()) == 1 {
				deadEnds = append(deadEnds, grid[row][column])
			}
		}
	}

	return deadEnds
}

func bestBraidCandidate(cells []cell.Cell) cell.Cell {
	for c := 0; c < len(cells); c++ {
		if len(cells[c].Links()) == 1 {
			return cells[c]
		}
	}

	return cells[rand.Intn(len(cells))]
}

func Braid(sh ShapeHolder, percentBraided float32) {
	deadEnds := deadEnds(sh)
	shuffle(deadEnds)

	if percentBraided > 1 {
		percentBraided = 1
	}
	if percentBraided <= 0 {
		return
	}

	// It will act similar to math.Floor
	toCull := int(float32(len(deadEnds)) * percentBraided)

	for c := 0; c < toCull; c++ {

		if len(deadEnds[c].Links()) != 1 {
			continue
		}

		neighbors := deadEnds[c].Neighbors()

		// maybe not needed
		if len(neighbors) == 0 {
			continue
		}

		deadEnds[c].Link(bestBraidCandidate(neighbors))

	}
}
