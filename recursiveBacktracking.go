package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

func RecursiveBacktracking(sh ShapeHolder) {
	stack := cell.CreateCellStack()
	stack.Push(sh.getShape().randomCell())

	for stack.Length() > 0 {
		current := stack.Peek()

		ns := current.Neighbors()
		neighbors := []cell.Cell{}

		for i := 0; i < len(ns); i++ {
			if len(ns[i].Links()) == 0 {
				neighbors = append(neighbors, ns[i])
			}
		}

		if len(neighbors) == 0 {
			stack.Pop()
		} else {
			neighbor := neighbors[rand.Intn(len(neighbors))]
			current.Link(neighbor)
			stack.Push(neighbor)
		}
	}

}
