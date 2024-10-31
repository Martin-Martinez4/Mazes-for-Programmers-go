package grid

import (
	"image"
	"image/color"
	"math"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
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
		Size:    rows * columns,
	}}

	prepareHexGrid(hg.getShape())
	configureHexCells(hg)

	return hg
}

func (hg *HexGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (hg *HexGrid) getShape() *Shape {
	return hg.Shape
}

func prepareHexGrid(g *Shape) {
	grid := make([][]cell.Cell, g.Rows())

	for row := 0; row < g.Rows(); row++ {
		grid[row] = make([]cell.Cell, g.Columns())

		for column := 0; column < g.Columns(); column++ {
			grid[row][column] = cell.CreateHexCell(row, column)
		}
	}

	g.Grid = grid
}

func configureHexCells(hg *HexGrid) {

	grid := hg.getShape().Grid

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

			if hg.GetCell(northDiagonal, column-1) == nil {

				c.NorthWest = nil
			} else {
				c.NorthWest = hg.GetCell(northDiagonal, column-1).(*cell.HexCell)
			}

			if hg.GetCell(row-1, column) == nil {

				c.North = nil
			} else {
				c.North = hg.GetCell(row-1, column).(*cell.HexCell)
			}

			if hg.GetCell(northDiagonal, column+1) == nil {

				c.NorthEast = nil
			} else {

				c.NorthEast = hg.GetCell(northDiagonal, column+1).(*cell.HexCell)
			}

			if hg.GetCell(southDiagonal, column-1) == nil {

				c.SouthWest = nil
			} else {

				c.SouthWest = hg.GetCell(southDiagonal, column-1).(*cell.HexCell)
			}

			if hg.GetCell(row+1, column) == nil {

				c.South = nil
			} else {

				c.South = hg.GetCell(row+1, column).(*cell.HexCell)
			}

			if hg.GetCell(southDiagonal, column+1) == nil {

				c.SouthEast = nil
			} else {

				c.SouthEast = hg.GetCell(southDiagonal, column+1).(*cell.HexCell)
			}

		}
	}
}

// this is going to need to be reworked
func (hg *HexGrid) ToPNG(filepath string, size int) {
	shape := hg.getShape()

	aSize := size / 2.0
	bSize := (float64(size) * math.Sqrt(3) / 2.0)

	// width := size * 2
	height := bSize * 2

	imgWidth := int((float64(3*aSize*shape.Columns()+aSize) + .5))
	imgHeight := int(height*float64(hg.Rows()) + bSize - .5)

	// background := imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}
	// wall := imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+1, imgHeight+1))

	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, white)
		}
	}

	// const BACKGROUND = 1
	// for mode := 0; mode < 2; mode++ {

	for row := 0; row < len(hg.getShape().Grid); row++ {
		for column := 0; column < len(hg.getShape().Grid[row]); column++ {
			cell := hg.GetCell(row, column).(*cell.HexCell)

			cx := size + 3*column*aSize
			cy := bSize + float64(row)*height

			if column%2 != 0 {
				cy += bSize
			}

			// f/n -> far/near
			// n/s/e/w -> north/south/east/west
			xFW := cx - size
			xNW := cx - aSize
			xNE := cx + aSize
			xFE := cx + size

			// m -> middle
			yN := int(cy - bSize)
			yM := int(cy)
			yS := int(cy + bSize)

			// if mode == BACKGROUND {
			// 	color := pixels[row][column]

			// 	points := [][]int{
			// 		{xFW, yM},
			// 		{xNW, yN},
			// 		{xNE, yN},
			// 		{xFE, yM},
			// 		{xNE, yS},
			// 		{xNW, yS},
			// 	}

			// 	// draw line with sliding window?
			// 	for i := 0;
			// } else {

			if cell.SouthWest == nil {
				draw.StraightLine2(xFW, yM, xNW, yS, img, black)
			}
			if cell.NorthWest == nil {
				draw.StraightLine2(xFW, yM, xNW, yN, img, black)
			}
			if cell.North == nil {

				draw.StraightLine2(xNW, yN, xNE, yN, img, black)
			}
			if !cell.IsLinked(cell.NorthEast) {
				draw.StraightLine2(xNE, yN, xFE, yM, img, black)
			}
			if !cell.IsLinked(cell.SouthEast) {
				draw.StraightLine2(xFE, yM, xNE, yS, img, black)
			}
			if !cell.IsLinked(cell.South) {
				draw.StraightLine2(xNE, yS, xNW, yS, img, black)
			}
			// }

			// }
		}

	}
	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)

}
