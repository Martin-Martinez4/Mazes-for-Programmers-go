package grid

import (
	"image"
	"image/color"
	"math"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	drw "github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

//	only algorithms like
//
// Wilson's, Recursive Backtracker, and Aldous-Broder will work with the library in its current state.

type PolarGrid struct {
	*Shape
}

func CreatePolarGrid(rows int) *PolarGrid {
	shape := &Shape{
		rows:    rows,
		columns: 1,
		Size:    rows * 1,
	}
	pg := &PolarGrid{Shape: shape}

	preparePolarGrid(pg)
	configurePolarCells(pg)

	return pg
}

func (pg *PolarGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (pg *PolarGrid) getShape() *Shape {
	return pg.Shape
}

func (pg *PolarGrid) GetCell(row int, column int) cell.Cell {
	grid := pg.getShape().Grid
	if row >= 0 && row <= len(grid)-1 {
		c := column % len(grid[row])
		if c < 0 {
			c = c + len(grid[row])
		}
		return grid[row][c]
	}
	return nil
}

// this is going to need to be reworked
func (pg *PolarGrid) ToPNG(filepath string, cellSize int) {
	shape := pg.getShape()
	grid := shape.Grid

	imgSize := 2 * pg.Rows() * cellSize

	// background := imagehandling.Pixel{R: 255, G: 255, B: 255, A: 255}
	wall := imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgSize+1, imgSize+1))
	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imgSize+1; x++ {
		for y := 0; y < imgSize+1; y++ {
			img.Set(x, y, white)
		}
	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgSize+1, imgSize+1)

	center := imgSize / 2

	// replace this
	for row := 1; row < len(pg.getShape().Grid); row++ {
		for column := 0; column < len(pg.getShape().Grid[row]); column++ {

			// c := grid[row][column]
			c := pg.GetCell(row, column)

			c2, ok := c.(*cell.PolarCell)
			if !ok {
				return
			}

			theta := 2 * math.Pi / float64(len(grid[c.Row()]))
			innerRadius := float64(c.Row() * cellSize)
			outerRadius := float64((c.Row() + 1) * cellSize)
			thetaCcw := float64(c.Column()) * theta
			thetaCw := float64(c.Column()+1) * theta

			ax := center + int(innerRadius*math.Cos(thetaCcw))
			ay := center + int(innerRadius*math.Sin(thetaCcw))
			bx := center + int(outerRadius*math.Cos(thetaCcw))
			by := center + int(outerRadius*math.Sin(thetaCcw))

			cx := center + int(innerRadius*math.Cos(thetaCw))
			cy := center + int(innerRadius*math.Sin(thetaCw))
			dx := center + int(outerRadius*math.Cos(thetaCw))
			dy := center + int(outerRadius*math.Sin(thetaCw))

			if !c2.IsLinked(c2.Inward) {
				drw.StraightLine(ax, ay, cx, cy, pixels, wall)
			}
			if !c2.IsLinked(c2.Cw) {
				drw.StraightLine(cx, cy, dx, dy, pixels, wall)
			}
			if c2.Row() == len(pg.getShape().Grid)-1 {
				drw.StraightLine(bx, by, dx, dy, pixels, wall)
			}

		}
	}

	imagehandling.WritePNGFromPixels(filepath, pixels)

}

func preparePolarGrid(pg *PolarGrid) {

	rs := pg.Rows()

	rows := make([][]cell.Cell, rs)

	rowHeight := 1.0 / float64(rs)

	pcell := cell.CreatePolarCell(0, 0)
	rows[0] = []cell.Cell{pcell}

	size := 0

	for row := 1; row < rs; row++ {

		radius := float64(row) / float64(rs)
		circumference := 2 * math.Pi * radius

		previousCount := len(rows[row-1])
		estimatedCellWidth := circumference / float64(previousCount)
		ratio := estimatedCellWidth / float64(rowHeight)

		cells := int(previousCount * int(math.Round(ratio)))
		size += cells
		tmp := make([]cell.Cell, cells)
		for i := 0; i < cells; i++ {

			tmp[i] = cell.CreatePolarCell(row, i)
		}

		rows[row] = tmp
	}

	pg.getShape().Grid = rows
	pg.Shape.Size = size

}

func configurePolarCells(pg *PolarGrid) {
	grid := pg.getShape().Grid

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {

			c := pg.GetCell(row, column).(*cell.PolarCell)
			ro, col := c.Row(), c.Column()

			if ro > 0 {

				c.Ccw = pg.GetCell(ro, column-1).(*cell.PolarCell)

				c.Cw = pg.GetCell(ro, column+1).(*cell.PolarCell)

				ratio := float64(len(grid[ro])) / float64(len(grid[ro-1]))
				parent := grid[ro-1][int(float64(col)/ratio)].(*cell.PolarCell)
				parent.Outward = append(parent.Outward, c)
				c.Inward = parent
			}
		}
	}

}
