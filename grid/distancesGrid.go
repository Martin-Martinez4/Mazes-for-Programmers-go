package grid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

const STARTPOINT = 64

type DistancesGrid struct {
	*cell.Distances
	Shape *Shape
}

func CreateDistancesGrid(rows, columns int) *DistancesGrid {

	shape := CreateShape(rows, columns)

	return &DistancesGrid{Shape: shape}
}

func (dg *DistancesGrid) setDistancesTo(row, column int) {
	dg.Distances = dg.Shape.Grid[row][column].Distances()
}

func (dg *DistancesGrid) ContentsOf(cell cell.Cell) string {
	if dg.Distances.MaxDistance == 0 {
		_, val := dg.Distances.Max()
		dg.Distances.MaxDistance = val
	}

	maxSpaces := DigitsInInt(dg.Distances.MaxDistance)
	spacesHave := DigitsInInt(dg.Distances.Cells[cell])

	spaces := (maxSpaces - spacesHave)

	if dg.Distances == nil {
		panic("distances were not initialized")
	} else {

		trail := strings.Repeat(" ", spaces)
		return fmt.Sprint(trail + strconv.Itoa(dg.Distances.Cells[cell]))

	}
}

func (dg *DistancesGrid) GetShape() *Shape {
	return dg.Shape
}

func (dg *DistancesGrid) ToPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(dg, cellSize, cellSize)
	withWalls := drawWalls(dg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)
}

func DigitsInInt(num int) int {
	if num <= 9 {
		return 1
	}
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}
