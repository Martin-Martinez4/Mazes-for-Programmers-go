package main

import (
	"image"
	"image/color"
	"math"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
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

	imgSize := 2 * shape.rows * cellSize

	// background := imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}
	// wall := imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgSize+1, imgSize+1))
	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imgSize+1; x++ {
		for y := 0; y < imgSize+1; y++ {
			img.Set(x, y, white)
		}
	}

	// pixels := imagehandling.PNGDataToPixelSlice(img, imgSize+1, imgSize+1)

	// center := imgSize / 2

	// for row := 1; row < len(pg.getShape().grid); row++ {
	// 	for column := 0; column < len(pg.getShape().grid[row]); column++ {

	// 		// c := grid[row][column]
	// 		c := pg.GetCell(row, column)

	// 		c2, ok := c.(*cell.PolarCell)
	// 		if !ok {
	// 			return
	// 		}

	// 		theta := 2 * math.Pi / float64(len(grid[c.Row()]))
	// 		innerRadius := float64(c.Row() * cellSize)
	// 		outerRadius := float64((c.Row() + 1) * cellSize)
	// 		thetaCcw := float64(c.Column()) * theta
	// 		thetaCw := float64(c.Column()+1) * theta

	// 		ax := center + int(innerRadius*math.Cos(thetaCcw))
	// 		ay := center + int(innerRadius*math.Sin(thetaCcw))
	// 		bx := center + int(outerRadius*math.Cos(thetaCcw))
	// 		by := center + int(outerRadius*math.Sin(thetaCcw))

	// 		cx := center + int(innerRadius*math.Cos(thetaCw))
	// 		cy := center + int(innerRadius*math.Sin(thetaCw))
	// 		dx := center + int(outerRadius*math.Cos(thetaCw))
	// 		dy := center + int(outerRadius*math.Sin(thetaCw))

	// 		if !c2.IsLinked(c2.Inward) {
	// 			drw.StraightLine(ax, ay, cx, cy, pixels, wall)
	// 		}
	// 		if !c2.IsLinked(c2.Cw) {
	// 			drw.StraightLine(cx, cy, dx, dy, pixels, wall)
	// 		}
	// 		if c2.Row() == len(pg.getShape().grid)-1 {
	// 			drw.StraightLine(bx, by, dx, dy, pixels, wall)
	// 		}

	// 	}
	// }

	// imagehandling.WritePNGFromPixels(filepath, pixels)

}
