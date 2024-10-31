package grid

import (
	"image"
	"image/color"
	"math/rand"

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
		rows:    rows,
		columns: columns,
		Size:    rows * columns,
	},
		Levels: levels,
	}

	prepareGrid3D(g3d)
	configureCells3D(g3d)

	return g3d
}

func (g3d *Grid3D) ContentsOf(c cell.Cell) string {
	return " "
}

func (g3d *Grid3D) GetShape() *Shape {
	return g3d.Shape
}

func (g3d *Grid3D) GetCell(level, row, column int) cell.Cell {
	g := g3d.Grid
	if (level >= 0 && level < len(g)) &&
		(row >= 0 && row < len(g[level])) &&
		(column >= 0 && column < len(g[level][row])) {
		return g[level][row][column]
	}
	return nil
}

func (g3d *Grid3D) RandomCell() cell.Cell {

	g := g3d.Grid
	// random int [0. rows)
	// and.Intn(max-min) + min, but min is 0
	randLevel := rand.Intn(g3d.Levels)
	randRow := rand.Intn(len(g[randLevel]))
	randColumn := rand.Intn(len(g[randLevel][randRow]))

	for g3d.GetCell(randLevel, randRow, randColumn) == nil {

		randLevel = rand.Intn(g3d.Levels)
		randRow = rand.Intn(len(g[randLevel]))
		randColumn = rand.Intn(len(g[randLevel][randRow]))

	}

	return g3d.GetCell(randLevel, randRow, randColumn)
}

func prepareGrid3D(g3d *Grid3D) {
	levels := g3d.Levels
	g := make([][][]cell.Cell, levels)

	for lv := 0; lv < levels; lv++ {
		g[lv] = make([][]cell.Cell, g3d.Rows())
		for row := 0; row < g3d.Rows(); row++ {

			g[lv][row] = make([]cell.Cell, g3d.Columns())

			for column := 0; column < g3d.Columns(); column++ {
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
		for row := 0; row < len(grid[lv]); row++ {
			for column := 0; column < len(grid[lv][row]); column++ {
				c, ok := grid[lv][row][column].(*cell.Cell3D)
				if !ok {
					panic("Could not cast to cell.Cell3D")
				}

				if g3d.GetCell(lv, row-1, column) != nil {

					c.North = g3d.GetCell(lv, row-1, column).(*cell.Cell3D)
				} else {
					c.North = nil
				}

				if g3d.GetCell(lv, row+1, column) != nil {

					c.South = g3d.GetCell(lv, row+1, column).(*cell.Cell3D)
				} else {
					c.South = nil
				}

				if g3d.GetCell(lv, row, column-1) != nil {

					c.West = g3d.GetCell(lv, row, column-1).(*cell.Cell3D)
				} else {
					c.West = nil
				}

				if g3d.GetCell(lv, row, column+1) != nil {

					c.East = g3d.GetCell(lv, row, column+1).(*cell.Cell3D)
				} else {
					c.East = nil
				}
				if g3d.GetCell(lv-1, row, column) != nil {

					c.Down = g3d.GetCell(lv-1, row, column).(*cell.Cell3D)
				} else {
					c.Down = nil
				}
				if g3d.GetCell(lv+1, row, column) != nil {

					c.Up = g3d.GetCell(lv+1, row, column).(*cell.Cell3D)
				} else {
					c.Up = nil
				}
			}
		}
	}
}

// this is going to need to be reworked
func (g3d *Grid3D) ToPNG(filepath string, size, margin int, inset float32) {
	grid := g3d.Grid

	insetInt := int(float32(size) * inset)

	gridWidth := size * g3d.Columns()
	gridHeight := size * g3d.Rows()

	imgWidth := gridWidth*g3d.Levels + (g3d.Levels-1)*margin
	imgHeight := gridHeight

	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	arrow := color.RGBA{R: 100, G: 20, B: 20, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+2, imgHeight+2))

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, white)
		}
	}

	for lv := 0; lv < g3d.Levels; lv++ {
		for row := 0; row < len(grid[lv]); row++ {
			for column := 0; column < len(grid[lv][row]); column++ {
				c := g3d.GetCell(lv, row, column).(*cell.Cell3D)

				x := lv*(gridWidth+margin) + column*size
				y := row * size

				if inset > 0 {
					withInset3D(img, c, size, x, y, insetInt, black)
				} else {

					noInset3D(img, c, size, x, y, black)
				}

				midX := x + size/2
				midY := y + size/2

				if c.IsLinked(c.Down) {
					draw.StraightLine2(midX-3, midY, midX-1, midY+2, img, arrow)
					draw.StraightLine2(midX-3, midY, midX-1, midY-2, img, arrow)
				}

				if c.IsLinked(c.Up) {
					draw.StraightLine2(midX+3, midY, midX+1, midY+2, img, arrow)
					draw.StraightLine2(midX+3, midY, midX+1, midY-2, img, arrow)
				}

			}

		}
	}
	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)

}

func noInset3D(img *image.RGBA, c *cell.Cell3D, cellSize, x, y int, wall color.RGBA) {
	x1, y1 := x, y
	x2 := x1 + cellSize
	y2 := y1 + cellSize

	if c.North == nil {

		draw.StraightLine2(x1, y1, x2, y1, img, wall)
	}

	if c.West == nil {

		draw.StraightLine2(x1, y1, x1, y2, img, wall)
	}

	if !c.IsLinked(c.East) {

		draw.StraightLine2(x2, y1, x2, y2, img, wall)
	}
	if !c.IsLinked(c.South) {

		draw.StraightLine2(x1, y2, x2, y2, img, wall)
	}

}

func withInset3D(img *image.RGBA, c *cell.Cell3D, cellSize, x, y, inset int, wall color.RGBA) {
	x1, x2, x3, x4, y1, y2, y3, y4 := cellCoordsWithInset(x, y, cellSize, inset)

	if c.IsLinked(c.North) {
		draw.StraightLine2(x2, y1, x2, y2, img, wall)
		draw.StraightLine2(x3, y1, x3, y2, img, wall)
	} else {
		draw.StraightLine2(x2, y2, x3, y2, img, wall)

	}

	if c.IsLinked(c.South) {
		draw.StraightLine2(x2, y3, x2, y4, img, wall)
		draw.StraightLine2(x3, y3, x3, y4, img, wall)
	} else {
		draw.StraightLine2(x2, y3, x3, y3, img, wall)

	}

	if c.IsLinked(c.West) {
		draw.StraightLine2(x1, y2, x2, y2, img, wall)
		draw.StraightLine2(x1, y3, x2, y3, img, wall)
	} else {
		draw.StraightLine2(x2, y2, x2, y3, img, wall)

	}

	if c.IsLinked(c.East) {
		draw.StraightLine2(x3, y2, x4, y2, img, wall)
		draw.StraightLine2(x3, y3, x4, y3, img, wall)
	} else {
		draw.StraightLine2(x3, y2, x3, y3, img, wall)

	}
}
