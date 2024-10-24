package main

import (
	"fmt"
	"math/rand"
)

func BinaryTree(sh ShapeHolder) {

	grid := sh.getShape()

	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			cell := grid.grid[row][column]
			c2, ok := cell.(*BaseCell)
			if !ok {
				return
			}

			neighbors := []Cell{}

			if c2.north != nil {
				neighbors = append(neighbors, c2.north)
			}

			if c2.east != nil {
				neighbors = append(neighbors, c2.east)
			}

			if len(neighbors) == 0 {
				continue
			} else {
				randIndex := rand.Intn(len(neighbors))
				c2.Link(neighbors[randIndex])
			}

		}
	}
}

func testa(a [][]Cell) {
	fmt.Println("test a")
}
