package grid

import (
	"image"
	"image/color"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

type PlainGrid struct {
	*Shape
}

func CreatePlainGrid(rows, columns int) *PlainGrid {
	shape := &Shape{
		rows:    rows,
		columns: columns,
		Size:    rows * columns,
	}

	prepareGrid(shape)
	configureCells(shape)

	return &PlainGrid{Shape: shape}
}

func (pg *PlainGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (pg *PlainGrid) GetShape() *Shape {
	return pg.Shape
}

func (pg *PlainGrid) ToPNG(filepath string, cellSize int) {
	pixs := PixelsFromShape(pg, cellSize, cellSize)
	withWalls := drawWalls(pg, pixs, 2)
	imagehandling.WritePNGFromPixels(filepath, withWalls)

}

func (pg *PlainGrid) Png(filepath string, cellSize int, inset float32) {

	grid := pg.GetShape().Grid

	imgWidth := cellSize * pg.Columns()
	imgHeight := cellSize * pg.Rows()

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
			c := pg.GetShape().GetCell(row, column)
			if c == nil {
				continue
			}

			cb := c.(*cell.BaseCell)

			x := column * cellSize
			y := row * cellSize

			if inset > 0 {
				withInset(img, cb, cellSize, x, y, insetInt, wall)
			} else {
				noInset(img, cb, cellSize, x, y, wall)
			}
		}
	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)
}

func noInset(img *image.RGBA, c *cell.BaseCell, cellSize, x, y int, wall color.RGBA) {
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

func withInset(img *image.RGBA, c *cell.BaseCell, cellSize, x, y, inset int, wall color.RGBA) {
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

func cellCoordsWithInset(x, y, cellSize, inset int) (x1, x2, x3, x4, y1, y2, y3, y4 int) {
	x_1, x_4 := x, x+cellSize
	x_2 := x_1 + inset
	x_3 := x_4 - inset

	y_1, y_4 := y, y+cellSize
	y_2 := y_1 + inset
	y_3 := y_4 - inset

	return x_1, x_2, x_3, x_4, y_1, y_2, y_3, y_4
}
