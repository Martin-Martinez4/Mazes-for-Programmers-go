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

type Grid3D struct {
	*Shape
	Grid   [][][]cell.Cell
	Levels int
}

func CreateGrid3D(levels, rows, columns int) *Grid3D {

	g3d := &Grid3D{Shape: &Shape{
		Rows:    rows,
		Columns: columns,
		Size:    rows * columns,
	},
		Levels: levels,
	}

	prepareHexGrid(g3d.getShape())
	configureCells3D(g3d)

	return g3d
}

func (g3d *Grid3D) ContentsOf(c cell.Cell) string {
	return " "
}

func (g3d *Grid3D) getShape() *Shape {
	return g3d.Shape
}

func prepareGrid3D(g3d *Grid3D) {
	levels := g3d.Levels
	g := make([][][]cell.Cell, levels)

	for lv := 0; lv < levels; lv++ {
		g[lv] = make([][]cell.Cell, g3d.Rows)
		for row := 0; row < g3d.Rows; row++ {

			g[lv][row] = make([]cell.Cell, g3d.Columns)

			for column := 0; column < g3d.Columns; column++ {
				g[lv][row][column] = cell.CreateCell3D(lv, row, column)
			}
		}
	}

	g3d.Grid = g
}

func configureCells3D(g3d *Grid3D) {

	grid := g3d.Grid
	levels := g3d.Levels

	for lv := 0; lv < levels; lv++ {
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid[row]); column++ {
				c := grid[lv][row][column].(*cell.Cell3D)

			}
		}
	}
}

// this is going to need to be reworked
func (g3d *Grid3D) ToPNG(filepath string, size int) {
	shape := g3d.getShape()

	aSize := size / 2.0
	bSize := (float64(size) * math.Sqrt(3) / 2.0)

	// width := size * 2
	height := bSize * 2

	imgWidth := int((float64(3*aSize*shape.Columns+aSize) + .5))
	imgHeight := int(height*float64(hg.Rows) + bSize - .5)

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
