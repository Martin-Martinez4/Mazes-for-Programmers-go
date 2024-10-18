package main

import "fmt"

func main() {
	fmt.Println("Sup")

	grid1 := CreateGrid(4, 4)

	Sidewinder(grid1)

	grid1.print()

	queue := CreateQueue(4)

	cell1 := CreateCell(1, 1)
	cell2 := CreateCell(1, 0)
	cell3 := CreateCell(0, 1)
	cell5 := CreateCell(1, 2)

	cell1.east = cell5
	cell1.west = cell2
	cell2.north = cell3

	queue.Push(cell1)

	cell := queue.Pop()
	for cell != nil {

		if cell.north != nil {
			queue.Push(cell.north)
		}
		if cell.south != nil {
			queue.Push(cell.south)
		}
		if cell.east != nil {
			queue.Push(cell.east)
		}
		if cell.west != nil {
			queue.Push(cell.west)
		}

		fmt.Println("Cell: ", cell.row, " ", cell.column)

		cell = queue.Pop()
	}
}
