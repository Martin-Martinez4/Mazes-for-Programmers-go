package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"

/*
	Create a draw line function
		- find slope between two points
		- use point slope formula in a loop to draw the pixels

		- if line is vertical (x_2 - x_1) is zero
			- Run Draw Veritcal line function
*/

func main() {

	mask := MaskFromPNG("./images/mask/circle_mask.png")
	maskgrid := CreateMaskGrid(mask)
	HuntAndKill(maskgrid)

	pixs := PixelsFromShape(maskgrid, 30, 30)
	withWalls := drawWalls(maskgrid, pixs, 2)

	imagehandling.WritePNGFromPixels("./images/masked_output", withWalls)

}
