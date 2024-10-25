package main

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
	size    int
	grid    [][]cell.Cell
}

func (g *Shape) GetCell(row int, column int) cell.Cell {
	if (row >= 0 && row < len(g.grid)) &&
		(column >= 0 && column < len(g.grid[row])) {
		return g.grid[row][column]
	}
	return nil
}

func CreateShape(rows, columns int) *Shape {
	shape := &Shape{
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}

	prepareGrid(shape)
	configureCells(shape)

	return shape
}

func prepareGrid(g *Shape) {
	grid := make([][]cell.Cell, g.rows)

	for row := 0; row < g.rows; row++ {
		grid[row] = make([]cell.Cell, g.columns)

		for column := 0; column < g.columns; column++ {
			grid[row][column] = cell.CreateBaseCell(row, column)
		}
	}

	g.grid = grid
}
func configureCells(g *Shape) {

	for row := 0; row < g.rows; row++ {
		for column := 0; column < g.columns; column++ {

			c := g.grid[row][column]

			if c != nil {
				c2, ok := c.(*cell.BaseCell)
				if !ok {
					return
				}

				if row > 0 {

					c2.North = g.grid[row-1][column].(*cell.BaseCell)
				}

				if row < g.rows-1 {

					c2.South = g.grid[row+1][column].(*cell.BaseCell)
				}

				if column > 0 {

					c2.West = g.grid[row][column-1].(*cell.BaseCell)
				}

				if column < g.columns-1 {

					c2.East = g.grid[row][column+1].(*cell.BaseCell)
				}
			}
		}
	}
}

func (g *Shape) randomCell() cell.Cell {

	// random int [0. rows)
	// and.Intn(max-min) + min, but min is 0
	randRow := rand.Intn(len(g.grid))
	randColumn := rand.Intn(len(g.grid[randRow]))

	for g.grid[randRow][randColumn] == nil {

		randRow = rand.Intn(len(g.grid))
		randColumn = rand.Intn(len(g.grid[randRow]))
	}

	return g.grid[randRow][randColumn]
}

// instead of taking in a grid take in an interface with get shape and cell contents
func print(sh ShapeHolder) {

	g := sh.getShape()

	// build top and bottom strings as the rows go
	var top strings.Builder
	var bottom strings.Builder

	// top
	fmt.Print("+")
	for columns := 0; columns < g.columns; columns++ {
		fmt.Print("---+")
	}
	fmt.Print("\n")

	for row := 0; row < g.rows; row++ {
		top.WriteString("|")
		bottom.WriteString("+")
		for cellIndex := 0; cellIndex < g.columns; cellIndex++ {

			var c *cell.BaseCell

			if g.grid[row][cellIndex] == nil {
				c = cell.CreateBaseCell(-1, -1)
			} else {
				c = g.grid[row][cellIndex].(*cell.BaseCell)
			}

			cellbody := fmt.Sprintf(" %s ", sh.ContentsOf(c))

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
