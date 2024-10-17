package main

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func CreateDistances(root *Cell) *Distances {

	return &Distances{
		root:  root,
		cells: map[*Cell]int{root: 0},
	}
}
