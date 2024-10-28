package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

// Random walk with loop erasing
func Wilsons(sh grid.ShapeHolder) {
	unvisited := make([]cell.Cell, sh.GetShape().Size)

	grid := sh.GetShape().Grid
	rows := sh.GetShape().Rows
	columns := sh.GetShape().Columns

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			unvisited[(row*columns)+column] = grid[row][column]
		}
	}

	randIndex := rand.Intn(len(unvisited))
	unvisited = append(unvisited[:randIndex], unvisited[randIndex+1:]...)

	for len(unvisited) > 0 {
		randIndex := rand.Intn(len(unvisited))
		c := unvisited[randIndex]

		path := []cell.Cell{c}

		for includesCell(c, unvisited) {
			neighbors := c.Neighbors()
			if len(neighbors) <= 0 {
				return
			}
			c = neighbors[rand.Intn(len(neighbors))]
			position := indexOf(c, path)

			if position != -1 {
				path = path[0 : position+1]
			} else {
				path = append(path, c)

			}
		}

		for index := 0; index < len(path)-1; index++ {
			path[index].Link(path[index+1])
			index2 := indexOf(path[index], unvisited)
			if index2 == -1 {
				panic("cell not found in path")
			}
			unvisited = append(unvisited[:index2], unvisited[index2+1:]...)
		}
	}
}

func includesCell(c cell.Cell, cells []cell.Cell) bool {

	for i := 0; i < len(cells); i++ {
		if cells[i] == c {
			return true
		}
	}

	return false
}

func indexOf(c cell.Cell, cells []cell.Cell) int {

	for i := 0; i < len(cells); i++ {
		if cells[i] == c {
			return i
		}
	}
	return -1
}
