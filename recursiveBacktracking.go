package main

import "math/rand"

func RecursiveBacktracking(sh ShapeHolder) {
	stack := CreateCellStack()
	stack.Push(sh.getShape().randomCell())

	for stack.Length() > 0 {
		current := stack.Peek()

		ns := current.neighbors()
		neighbors := []*Cell{}

		for i := 0; i < len(ns); i++ {
			if len(ns[i].links) == 0 {
				neighbors = append(neighbors, ns[i])
			}
		}

		if len(neighbors) == 0 {
			stack.Pop()
		} else {
			neighbor := neighbors[rand.Intn(len(neighbors))]
			current.link(neighbor)
			stack.Push(neighbor)
		}
	}

}
