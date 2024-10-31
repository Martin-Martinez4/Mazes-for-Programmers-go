package grid

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

// Need to make a Basic Grid or Basic Shape Struct

// There is an overall grid struct, that struct takes in an interface, that interface has the method that would be overloaded in oop
type Shape struct {
	rows    int
	columns int
	Size    int
	Grid    [][]cell.Cell
}

func (g *Shape) GetCell(row int, column int) cell.Cell {
	if (row >= 0 && row < len(g.Grid)) &&
		(column >= 0 && column < len(g.Grid[row])) {
		return g.Grid[row][column]
	}
	return nil
}

func (g *Shape) Rows() int {
	return g.rows
}

func (g *Shape) Columns() int {
	return g.columns
}

func (g *Shape) EachCell() <-chan cell.Cell {
	ch := make(chan cell.Cell, 1)
	go func() {
		for row := 0; row < len(g.Grid); row++ {
			for column := 0; column < len(g.Grid[row]); column++ {
				c := g.Grid[row][column]
				if c != nil {
					ch <- c
				}
			}
		}
		close(ch)
	}()
	return ch
}

func CreateShape(rows, columns int) *Shape {
	shape := &Shape{
		rows:    rows,
		columns: columns,
		Size:    rows * columns,
	}

	prepareGrid(shape)
	configureCells(shape)

	return shape
}

func prepareGrid(g *Shape) {
	grid := make([][]cell.Cell, g.Rows())

	for row := 0; row < g.Rows(); row++ {
		grid[row] = make([]cell.Cell, g.Columns())

		for column := 0; column < g.Columns(); column++ {
			grid[row][column] = cell.CreateBaseCell(row, column)
		}
	}

	g.Grid = grid
}
func configureCells(g *Shape) {

	for row := 0; row < g.Rows(); row++ {
		for column := 0; column < g.Columns(); column++ {

			c := g.Grid[row][column]

			if c != nil {
				c2, ok := c.(*cell.BaseCell)
				if !ok {
					return
				}

				if row > 0 {

					c := g.GetCell(row-1, column)
					if c == nil {

						c2.North = nil
					} else {
						c2.North = c.(*cell.BaseCell)
					}

					// c2.North = g.grid[row-1][column].(*cell.BaseCell)
				}

				if row < g.Rows()-1 {

					c := g.GetCell(row+1, column)
					if c == nil {

						c2.South = nil
					} else {
						c2.South = c.(*cell.BaseCell)
					}

					// c2.South = g.grid[row+1][column].(*cell.BaseCell)
				}

				if column > 0 {

					c := g.GetCell(row, column-1)
					if c == nil {

						c2.West = nil
					} else {
						c2.West = c.(*cell.BaseCell)
					}
				}

				if column < g.Columns()-1 {
					c := g.GetCell(row, column+1)
					if c == nil {

						c2.East = nil
					} else {
						c2.East = c.(*cell.BaseCell)
					}

				}
			}
		}
	}
}

func (g *Shape) RandomCell() cell.Cell {

	// random int [0. rows)
	// and.Intn(max-min) + min, but min is 0
	randRow := rand.Intn(len(g.Grid))
	randColumn := rand.Intn(len(g.Grid[randRow]))

	for g.Grid[randRow][randColumn] == nil {

		randRow = rand.Intn(len(g.Grid))
		randColumn = rand.Intn(len(g.Grid[randRow]))
	}

	return g.Grid[randRow][randColumn]
}

// instead of taking in a grid take in an interface with get shape and cell contents
func Print(sh ShapeHolder) {

	g := sh.GetShape()

	// build top and bottom strings as the rows go
	var top strings.Builder
	var bottom strings.Builder

	// top
	fmt.Print("+")
	for columns := 0; columns < g.Columns(); columns++ {
		fmt.Print("---+")
	}
	fmt.Print("\n")

	for row := 0; row < g.Rows(); row++ {
		top.WriteString("|")
		bottom.WriteString("+")
		for cellIndex := 0; cellIndex < g.Columns(); cellIndex++ {

			var c *cell.BaseCell

			if g.Grid[row][cellIndex] == nil {
				c = cell.CreateBaseCell(-1, -1)
			} else {
				c = g.Grid[row][cellIndex].(*cell.BaseCell)
			}

			cellbody := fmt.Sprintf(" %s ", " ")

			var eastBoundary string
			if c.IsLinked(c.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}

			// fmt.Print(cellbody, eastBoundary)
			top.WriteString(cellbody)
			top.WriteString(eastBoundary)

			var southBoundary string
			if c.IsLinked(c.South) {
				// These need to increase with number size increase
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}

			corner := "+"

			bottom.WriteString(southBoundary)
			bottom.WriteString(corner)
		}

		fmt.Print(top.String(), "\n")
		fmt.Print(bottom.String(), "\n")

		top.Reset()
		bottom.Reset()
	}
}

func PrintDistanceGrid(dg *DistancesGrid) {

	g := dg.GetShape()
	_, max := dg.Distances.Max()
	maxSpaces := DigitsInInt(max)
	trail := strings.Repeat(" ", maxSpaces)
	dashes := strings.Repeat("-", maxSpaces)

	// build top and bottom strings as the rows go
	var top strings.Builder
	var bottom strings.Builder

	// top
	fmt.Print("+")
	for columns := 0; columns < g.Columns(); columns++ {
		fmt.Print(dashes + "--+")
	}
	fmt.Print("\n")

	for row := 0; row < g.Rows(); row++ {
		top.WriteString("|")
		bottom.WriteString("+")
		for cellIndex := 0; cellIndex < g.Columns(); cellIndex++ {

			var c *cell.BaseCell

			if g.Grid[row][cellIndex] == nil {
				c = cell.CreateBaseCell(-1, -1)
			} else {
				c = g.Grid[row][cellIndex].(*cell.BaseCell)
			}

			cellbody := fmt.Sprintf(" %s ", dg.ContentsOf(c))

			var eastBoundary string
			if c.IsLinked(c.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}

			// fmt.Print(cellbody, eastBoundary)
			top.WriteString(cellbody)
			top.WriteString(eastBoundary)

			var southBoundary string
			if c.IsLinked(c.South) {
				// These need to increase with number size increase
				southBoundary = trail + "  "
			} else {
				southBoundary = dashes + "--"
			}

			corner := "+"

			bottom.WriteString(southBoundary)
			bottom.WriteString(corner)
		}

		fmt.Print(top.String(), "\n")
		fmt.Print(bottom.String(), "\n")

		top.Reset()
		bottom.Reset()
	}
}
