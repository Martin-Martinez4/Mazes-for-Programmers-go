package main

func main() {

	// Special kinds of grids should have a normal grid as a field
	// shared functionality between all grids should be pointer functions on the normal grid field (maybe call it the shape?)
	// What would be overloaded methods in OOP, should be part of an interface.
	// Unqiue fields go on the special grid struct.

	grid1 := CreateDistancesGrid(5, 5)

	distanceGrid := grid1.SpecialGrid.(*DistancesGrid)

	Sidewinder(grid1)

	distanceGrid.Distances = grid1.grid[0][0].distances()

	grid1.print()

}
