package main

/*
	Create a draw line function
		- find slope between two points
		- use point slope formula in a loop to draw the pixels

		- if line is vertical (x_2 - x_1) is zero
			- Run Draw Veritcal line function
*/

func main() {

	/*
		The whole Shape thing might be obsolete.  Will look into removing.
	*/

	// Plain grid to PNG
	// grid1 := CreatePlainGrid(40, 30)
	// RecursiveBacktracking(grid1)
	// Braid(grid1, 1)
	// grid1.toPNG("./images/test_output/toPNG_Plain", 40)
	// print(grid1)

	// Masked grid toPNG
	// mask := MaskFromPNG("./images/mask/circle_mask.png")
	// maskgrid := CreateMaskGrid(mask)
	// HuntAndKill(maskgrid)
	// Braid(maskgrid, 1)
	// maskgrid.toPNG("./images/test_output/toPNG_Masked", 45)

	// pGrid := CreatePolarGrid(20)
	// AldousBroder(pGrid)
	// Braid(pGrid, 1)
	// pGrid.toPNG("./images/test_output/toPNG_Polar_2", 40)

	// hGrid := CreateHexGrid(10, 12)
	// HuntAndKill(hGrid)
	// hGrid.toPNG("./images/test_output/toPNG_Hex", 100)

	// tGrid := CreateTriangleGrid(8, 8)
	// RecursiveBacktracking(tGrid)
	// tGrid.toPNG("./images/test_output/toPNG_Tri", 300)

}
