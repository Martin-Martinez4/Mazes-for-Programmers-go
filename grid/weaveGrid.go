package grid

import (
	"image"
	"image/color"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

type WeaveGrid struct {
	*Shape
	UnderCells []*OverUnderCell
}

func CreateWeaveGrid(rows, columns int) *WeaveGrid {

	wg := &WeaveGrid{Shape: &Shape{
		rows:    rows,
		columns: columns,
		Size:    rows * columns,
	}}

	prepareWeaveGrid(wg.Shape, wg)
	configureWeaveCells(wg.Shape)

	return wg
}

func (pg *WeaveGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (pg *WeaveGrid) GetShape() *Shape {
	return pg.Shape
}

func (wg *WeaveGrid) ToPNG(filepath string, cellSize int) {
	grid := wg.GetShape().Grid
	var inset float32 = 0.1

	imgWidth := cellSize * wg.Columns()
	imgHeight := cellSize * wg.Rows()

	insetInt := int(float32(cellSize) * inset)

	wall := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+1, imgHeight+1))
	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imgWidth+1; x++ {
		for y := 0; y < imgHeight+1; y++ {
			img.Set(x, y, white)
		}
	}

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			c := wg.GetShape().GetCell(row, column)
			if c == nil {
				continue
			}

			cb := c.(*OverUnderCell)

			x := column * cellSize
			y := row * cellSize

			wg.withInset(img, cb, cellSize, x, y, insetInt, wall)

		}
	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)
}

func (wg *WeaveGrid) withInset(img *image.RGBA, c *OverUnderCell, cellSize, x, y, inset int, wall color.RGBA) {
	x1, x2, x3, x4, y1, y2, y3, y4 := cellCoordsWithInset(x, y, cellSize, inset)

	if !c.IsUnder {
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
	} else {

		color := color.RGBA{R: 255, G: 0, B: 0, A: 255}

		if c.VerticalPassage() {
			draw.StraightLine2(x2, y1, x2, y2, img, color)
			draw.StraightLine2(x3, y1, x3, y2, img, color)
			draw.StraightLine2(x2, y3, x2, y4, img, color)
			draw.StraightLine2(x3, y3, x3, y4, img, color)
		} else {
			draw.StraightLine2(x1, y2, x2, y2, img, color)
			draw.StraightLine2(x1, y3, x2, y3, img, color)
			draw.StraightLine2(x3, y2, x4, y2, img, color)
			draw.StraightLine2(x3, y3, x4, y3, img, color)
		}
	}

}

func configureWeaveCells(g *Shape) {

	// replace this later
	for row := 0; row < g.Rows(); row++ {
		for column := 0; column < g.Columns(); column++ {

			c := g.Grid[row][column]

			if c != nil {
				c2, ok := c.(*OverUnderCell)
				if !ok {
					return
				}

				if row > 0 {

					c := g.GetCell(row-1, column)
					if c == nil {

						c2.North = nil
					} else {
						c2.North = c.(*OverUnderCell)
					}

					// c2.North = g.grid[row-1][column].(*OverUnderCell)
				}

				if row < g.Rows()-1 {

					c := g.GetCell(row+1, column)
					if c == nil {

						c2.South = nil
					} else {
						c2.South = c.(*OverUnderCell)
					}

					// c2.South = g.grid[row+1][column].(*OverUnderCell)
				}

				if column > 0 {

					c := g.GetCell(row, column-1)
					if c == nil {

						c2.West = nil
					} else {
						c2.West = c.(*OverUnderCell)
					}
				}

				if column < g.Columns()-1 {
					c := g.GetCell(row, column+1)
					if c == nil {

						c2.East = nil
					} else {
						c2.East = c.(*OverUnderCell)
					}

				}
			}
		}
	}
}

// It adds an under cell
func (wg *WeaveGrid) TunnelUnder(overCell *OverUnderCell) {
	uc := CreateUnderCell(overCell)
	wg.UnderCells = append(wg.UnderCells, uc)
}

func prepareWeaveGrid(shape *Shape, wg *WeaveGrid) {
	grid := make([][]cell.Cell, wg.Rows())

	for row := 0; row < wg.Rows(); row++ {
		grid[row] = make([]cell.Cell, wg.Columns())

		for column := 0; column < wg.Columns(); column++ {
			grid[row][column] = CreateOverCell(row, column, wg)
		}
	}

	shape.Grid = grid
}
