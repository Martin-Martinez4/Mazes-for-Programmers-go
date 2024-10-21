package main

func main() {

	// grid1 := CreatePlainGrid(5, 5)
	// RecursiveBacktracking(grid1)
	// print(grid1)

	// grid1.setDistancesTo(0, 0)
	// grid1.Distances = grid1.shortestPath(grid1.getCell(7, 4))
	// print(grid1)

	// mask := CreateMask(5, 5)
	// mask.turnOff(0, 0)

	// maskedGrid := CreateMaskGrid(mask)
	// RecursiveBacktracking(maskedGrid)
	// print(maskedGrid)

	mask := MaskFromFile("./mask.txt")
	maskedGrid := CreateMaskGrid(mask)

	RecursiveBacktracking(maskedGrid)
	print(maskedGrid)

}
