package main

/*
	Create a draw line function
		- find slope between two points
		- use point slope formula in a loop to draw the pixels

		- if line is vertical (x_2 - x_1) is zero
			- Run Draw Veritcal line function
*/

func main() {

	// Plain grid to PNG
	// grid1 := CreatePlainGrid(40, 30)
	// RecursiveBacktracking(grid1)
	// grid1.toPNG("./images/test_output/toPNG_Plain", 40)

	// Masked grid toPNG
	// mask := MaskFromPNG("./images/mask/circle_mask.png")
	// maskgrid := CreateMaskGrid(mask)
	// HuntAndKill(maskgrid)
	// maskgrid.toPNG("./images/test_output/toPNG_Masked", 45)

	pGrid := CreatePolarGrid(8, 20)
	RecursiveBacktracking(pGrid)
	pGrid.toPNG("./images/test_output/toPNG_Polar_1", 40)

}
