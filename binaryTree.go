package main

import (
	"fmt"
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

func BinaryTree(sh ShapeHolder) {

	grid := sh.getShape()

	for row := 0; row < grid.rows; row++ {

		for column := 0; column < grid.columns; column++ {
			c2, ok := (grid.grid[row][column]).(*cell.BaseCell)
			if !ok {
				return
			}

			neighbors := []cell.Cell{}

			if c2.North != nil {
				neighbors = append(neighbors, c2.North)
			}

			if c2.East != nil {
				neighbors = append(neighbors, c2.East)
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

func testa(a [][]cell.Cell) {
	fmt.Println("test a")
}
