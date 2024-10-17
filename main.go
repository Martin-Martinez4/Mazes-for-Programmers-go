package main

import "fmt"

func main() {
	fmt.Println("Sup")

	grid1 := CreateGrid(4, 4)

	BinaryTree(grid1)

	grid1.print()
}
