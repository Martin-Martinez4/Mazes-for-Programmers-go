package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const STARTPOINT = 64

// There is an overall grid struct, that struct takes in an interface, that interface has the method that would be overloaded in oop
type Shape struct {
	rows    int
	columns int
	size    int
	grid    [][]*Cell
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
	grid := make([][]*Cell, g.rows)

	for row := 0; row < g.rows; row++ {
		grid[row] = make([]*Cell, g.columns)

		for column := 0; column < g.columns; column++ {
			grid[row][column] = CreateCell(row, column)
		}
	}

	g.grid = grid
}
func configureCells(g *Shape) {

	for row := 0; row < g.rows; row++ {
		for column := 0; column < g.columns; column++ {
			cell := g.grid[row][column]

			if cell != nil {

				if row > 0 {

					cell.north = g.grid[row-1][column]
				}

				if row < g.columns-1 {

					cell.south = g.grid[row+1][column]
				}

				if column > 0 {

					cell.west = g.grid[row][column-1]
				}

				if column < g.columns-1 {

					cell.east = g.grid[row][column+1]
				}
			}
		}
	}
}

// func (g *Shape) randomCell() *Cell {

// 	// random int [0. rows)
// 	// and.Intn(max-min) + min, but min is 0
// 	randRow := rand.Intn(g.rows)
// 	randColumn := rand.Intn(g.columns)

//		return g.grid[randRow][randColumn]
//	}
func (g *Shape) randomCell() *Cell {

	// random int [0. rows)
	// and.Intn(max-min) + min, but min is 0
	randRow := rand.Intn(len(g.grid))
	randColumn := rand.Intn(len(g.grid[0]))

	for g.grid[randRow][randColumn] == nil {

		randRow = rand.Intn(len(g.grid))
		randColumn = rand.Intn(len(g.grid[0]))
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

			var cell *Cell

			if g.grid[row][cellIndex] == nil {
				cell = CreateCell(-1, -1)
			} else {
				cell = g.grid[row][cellIndex]
			}

			cellbody := fmt.Sprintf(" %s ", sh.ContentsOf(cell))

			var eastBoundary string
			if cell.isLinked(cell.east) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}

			// fmt.Print(cellbody, eastBoundary)
			top.WriteString(cellbody)
			top.WriteString(eastBoundary)

			var southBoundary string
			if cell.isLinked(cell.south) {
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
