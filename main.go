package main

import "fmt"

func main() {
	fmt.Println("Sup")

	grid1 := CreateGrid(4, 4)

	Sidewinder(grid1)

	grid1.print()
}
