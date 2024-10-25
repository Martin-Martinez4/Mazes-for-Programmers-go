package main

import (
	"image"
	"image/color"
	"math"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

//	only algorithms like
//
// Wilson's, Recursive Backtracker, and Aldous-Broder will work with the library in its current state.

type HexGrid struct {
	*Shape
}

func CreateHexGrid(rows, columns int) *HexGrid {

	hg := &HexGrid{Shape: &Shape{
		rows:    rows,
		columns: columns,
		size:    rows * columns,
	}}

	prepareGrid(hg.getShape())
	configureHexCells(hg)

	return hg
}

func (hg *HexGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (hg *HexGrid) getShape() *Shape {
	return hg.Shape
}

func configureHexCells(hg *HexGrid) {

	grid := hg.getShape().grid

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {

			var northDiagonal int
			var southDiagonal int

			if column%2 == 0 {
				northDiagonal = row - 1
				southDiagonal = row
			} else {
				northDiagonal = row
				southDiagonal = row + 1
			}

			c := hg.GetCell(row, column).(*cell.HexCell)

			c.NorthWest = hg.GetCell(northDiagonal, column-1).(*cell.HexCell)
			c.North = hg.GetCell(row-1, column).(*cell.HexCell)
			c.NorthEast = hg.GetCell(northDiagonal, column+1).(*cell.HexCell)

			c.SouthWest = hg.GetCell(southDiagonal, column-1).(*cell.HexCell)
			c.South = hg.GetCell(row+1, column).(*cell.HexCell)
			c.SouthEast = hg.GetCell(southDiagonal, column+1).(*cell.HexCell)

		}
	}
}

// this is going to need to be reworked
func (hg *HexGrid) toPNG(filepath string, cellSize int) {
	shape := hg.getShape()

	aSize := cellSize / 2.0
	bSize := float64(cellSize) * math.Sqrt(3) / 2.0

	width := cellSize * 2
	height := bSize * 2

	imgWidth := int(float64(3*aSize*shape.columns+aSize) + 0.5)
	imgHeight := int(height*float64(hg.rows) + bSize + 0.5)

	// background := imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}
	// wall := imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+1, imgHeight+1))
	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imgHeight; x++ {
		for y := 0; y < imgWidth; y++ {
			img.Set(x, y, white)
		}
	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	const BACKGROUND = 1
	for mode := 0; mode < 2; mode++ {

		for row := 1; row < len(hg.getShape().grid); row++ {
			for column := 0; column < len(hg.getShape().grid[row]); column++ {
				cx := cellSize + 3*column*aSize
				cy := bSize + float64(row)*height

				if column%2 != 0 {
					cy += bSize
				}

				// f/n -> far/near
				// n/s/e/w -> north/south/east/west
				xFW := cx - int(bSize)
				xNW := cx - aSize
				xNE := cx + aSize
				xFE := cx + cellSize

				// m -> middle
				yN := int(cy - bSize)
				yM := int(cy)
				yS := int(cy + bSize)

				if mode == BACKGROUND {
					color := pixels[row][column]

					points := [][]int{
						{xFW, yM},
						{xNW, yN},
						{xNE, yN},
						{xFE, yM},
						{xNE, yS},
						{xNW, yS},
					}

					// draw line with sliding window?
				} else {

				}

			}
		}
	}

	imagehandling.WritePNGFromPixels(filepath, pixels)

}
