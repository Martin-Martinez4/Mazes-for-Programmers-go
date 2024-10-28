package main

import (
	"fmt"
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func RecursiveBacktracking(sh grid.ShapeHolder) {
	stack := cell.CreateCellStack()
	stack.Push(sh.GetShape().RandomCell())
	for stack.Length() > 0 {
		current := stack.Peek()

		ns := current.Neighbors()
		neighbors := []cell.Cell{}

		for i := 0; i < len(ns); i++ {
			if len(ns[i].Links()) == 0 {
				neighbors = append(neighbors, ns[i])
			}
		}

		fmt.Println(len(neighbors))
		if len(neighbors) == 0 {
			stack.Pop()
		} else {
			neighbor := neighbors[rand.Intn(len(neighbors))]
			current.Link(neighbor)
			stack.Push(neighbor)
		}
	}

}
