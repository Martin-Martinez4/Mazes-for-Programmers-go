package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"

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

	// pg1 := CreatePlainGrid(5, 5)
	// RecursiveBacktracking(pg1)
	// Braid(pg1, .5)
	// pg1.Shape.GetCell(0, 1).SetWeight(3)
	// pg1.Shape.GetCell(2, 3).SetWeight(10)

	// dist := cell.WeightedShortestPath(pg1.Shape.GetCell(0, 0))
	// dg := &DistancesGrid{Distances: dist, Shape: pg1.Shape}
	// print(dg)

	// bg := grid.CreatePlainGrid(20, 20)
	// HuntAndKill(bg)
	// grid.Braid(bg, .2)
	// bg.Png("./images/test_output/inset_test", 100, 0.1)
	// bg.ToPNG("./images/test_output/no_inset_test", 100)

	wg := grid.CreateWeaveGrid(20, 20)
	RecursiveBacktracking(wg)
	wg.ToPNG("./images/test_output/weave_test", 100)
}
