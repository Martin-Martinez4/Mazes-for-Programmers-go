package main

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"

/*
	Mask from png
	- Create pixels from png
		- ReadPNG
		- PNGDataToPixelSlice
		- one pixel is one cell in [][]*Cell
		- if the pixel is black the Cell is nil
*/

func main() {

	mask := MaskFromPNG("./images/mask/circle_mask.png")
	maskgrid := CreateMaskGrid(mask)
	HuntAndKill(maskgrid)

	pixs := PixelsFromShape(maskgrid, 30, 30)
	withWalls := drawWalls(maskgrid, pixs, 2)

	imagehandling.WritePNGFromPixels("./images/masked_output", withWalls)

}
