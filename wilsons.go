package main

import (
	"math/rand"
)

// Random walk with loop erasing
func Wilsons(sh ShapeHolder) {
	unvisited := make([]Cell, sh.getShape().size)

	grid := sh.getShape().grid
	rows := sh.getShape().rows
	columns := sh.getShape().columns

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			unvisited[(row*columns)+column] = grid[row][column]
		}
	}

	randIndex := rand.Intn(len(unvisited))
	unvisited = append(unvisited[:randIndex], unvisited[randIndex+1:]...)

	for len(unvisited) > 0 {
		randIndex := rand.Intn(len(unvisited))
		cell := unvisited[randIndex]

		path := []Cell{cell}

		for includesCell(cell, unvisited) {
			neighbors := cell.Neighbors()
			cell = neighbors[rand.Intn(len(neighbors))]
			position := indexOf(cell, path)

			if position != -1 {
				path = path[0 : position+1]
			} else {
				path = append(path, cell)

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

func includesCell(cell Cell, cells []Cell) bool {

	for i := 0; i < len(cells); i++ {
		if cells[i] == cell {
			return true
		}
	}

	return false
}

func indexOf(cell Cell, cells []Cell) int {

	for i := 0; i < len(cells); i++ {
		if cells[i] == cell {
			return i
		}
	}
	return -1
}
