package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"

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

	// mask := MaskFromFile("./mask.txt")
	// maskedGrid := CreateMaskGrid(mask)

	// RecursiveBacktracking(maskedGrid)
	// print(maskedGrid)

	// img, w, h := imagehandling.ReadPNG("./images/test/do_not_delete.png")
	// pixels := imagehandling.PNGDataToPixelSlice(img, w, h)

	// imagehandling.WritePNGFromPixels("./images/test/copy", pixels)

	grid1 := CreatePlainGrid(1, 2)
	print(grid1)
	pixels := PixelsFromShape(grid1, 10, 10)

	imagehandling.WritePNGFromPixels("./images/worked", pixels)
}
