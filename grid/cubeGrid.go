package grid

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	d "github.com/Martin-Martinez4/Mazes-for-Programmers-go/draw"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

type CubeGrid struct {
	*Shape
	Grid  [][][]cell.Cell
	faces int
}

func CreateCubeGrid(size int) *CubeGrid {

	ccg := &CubeGrid{
		Shape: &Shape{
			rows:    size,
			columns: size,
			Size:    size * size,
		},
		faces: 6,
	}

	prepareCubeGrid(ccg)
	configureCubeCells(ccg)

	return ccg
}

func (ccg *CubeGrid) ContentsOf(c cell.Cell) string {
	return " "
}

func (ccg *CubeGrid) GetShape() *Shape {
	return ccg.Shape
}

// change this
func (g3d *CubeGrid) GetCell(face, row, column int) cell.Cell {

	if face < 0 || face >= 6 {
		return nil
	}
	f, r, c := g3d.wrap(face, row, column)
	return g3d.Grid[f][r][c]

}

func (g3d *CubeGrid) wrap(face, row, column int) (int, int, int) {
	n := g3d.Rows() - 1

	if row < 0 {
		switch face {
		case 0:
			return 4, column, 0
		case 1:
			return 4, n, column
		case 2:
			return 4, n - column, n
		case 3:
			return 4, 0, n - column
		case 4:
			return 3, 0, n - column
		case 5:
			return 1, n, column
		default:
			panic("Invalid number of faces row < 0")

		}
	} else if row >= n+1 {
		switch face {
		case 0:
			return 5, n - column, 0
		case 1:
			return 5, 0, column
		case 2:
			return 5, column, n
		case 3:
			return 5, n, n - column
		case 4:
			return 1, 0, column
		case 5:
			return 3, n, n - column
		default:
			panic("Invalid number of faces row >- n+1")

		}
	} else if column < 0 {
		switch face {
		case 0:
			return 3, row, n
		case 1:
			return 0, row, n
		case 2:
			return 1, row, n
		case 3:
			return 2, row, n
		case 4:
			return 0, 0, row
		case 5:
			return 0, n, n - row
		default:
			panic("Invalid number of faces column < 0")

		}
	} else if column >= n+1 {
		switch face {
		case 0:
			return 1, row, 0
		case 1:
			return 2, row, n
		case 2:
			return 3, row, n
		case 3:
			return 0, row, 0
		case 4:
			return 2, 0, n - row
		case 5:
			return 2, n, row
		default:
			panic("Invalid number of faces column >= n+1")

		}
	}

	return face, row, column
}

func (g3d *CubeGrid) RandomCell() cell.Cell {

	g := g3d.Grid
	// random int [0. rows)
	// and.Intn(max-min) + min, but min is 0
	randLevel := rand.Intn(g3d.faces)
	randRow := rand.Intn(len(g[randLevel]))
	randColumn := rand.Intn(len(g[randLevel][randRow]))

	for g3d.GetCell(randLevel, randRow, randColumn) == nil {

		randLevel = rand.Intn(g3d.faces)
		randRow = rand.Intn(len(g[randLevel]))
		randColumn = rand.Intn(len(g[randLevel][randRow]))

	}

	return g3d.GetCell(randLevel, randRow, randColumn)
}

func (g *CubeGrid) EachCell() <-chan cell.Cell {
	ch := make(chan cell.Cell, 1)
	go func() {
		for fc := 0; fc < len(g.Grid); fc++ {

			for row := 0; row < len(g.Grid[fc]); row++ {
				for column := 0; column < len(g.Grid[fc][row]); column++ {
					c := g.Grid[fc][row][column]
					if c != nil {
						ch <- c
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func prepareCubeGrid(ccg *CubeGrid) {

	numFaces := 6

	grid := make([][][]cell.Cell, numFaces)

	for f := 0; f < numFaces; f++ {

		grid[f] = make([][]cell.Cell, ccg.Rows())

		for row := 0; row < ccg.Rows(); row++ {
			grid[f][row] = make([]cell.Cell, ccg.Columns())

			for column := 0; column < ccg.Columns(); column++ {
				grid[f][row][column] = cell.CreateCubeCell(f, row, column)
			}
		}
	}

	ccg.Grid = grid
}

func configureCubeCells(ccg *CubeGrid) {
	for c := range ccg.EachCell() {
		c, ok := c.(*cell.CubeCell)
		if !ok {
			panic("Cast to cell>CubeCell failed")
		}

		face, row, column := c.Face(), c.Row(), c.Column()

		w := ccg.GetCell(face, row, column-1)
		if w != nil {
			c.West = w.(*cell.CubeCell)
		}

		e := ccg.GetCell(face, row, column+1)
		if e != nil {
			c.East = e.(*cell.CubeCell)
		}

		n := ccg.GetCell(face, row-1, column)
		if n != nil {
			c.North = n.(*cell.CubeCell)
		}

		s := ccg.GetCell(face, row+1, column)
		if s != nil {
			c.South = s.(*cell.CubeCell)
		}

	}
}

// this is going to need to be reworked
func (ccg *CubeGrid) ToPNG(filepath string, size int, inset float32) {

	insetInt := int(float32(size) * inset)

	faceWidth := size * ccg.Rows()
	faceHeight := size * ccg.Rows()

	imgWidth := 4 * faceWidth
	imgHeight := 3 * faceHeight

	offsets := [][]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {1, 0}, {1, 2}}

	// give pixels color
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	outline := color.RGBA{R: 50, G: 100, B: 50, A: 255}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth+1, imgHeight+1))
	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			img.Set(x, y, white)
		}
	}

	drawOutlines(img, faceWidth, faceHeight, outline)

	for c := range ccg.EachCell() {
		cc := c.(*cell.CubeCell)
		x := offsets[cc.Face()][0]*faceWidth + cc.Column()*size
		y := offsets[cc.Face()][1]*faceHeight + cc.Row()*size

		if inset > 0 {
			withInsetCube(img, cc, size, x, y, insetInt, black)
		} else {

			noInsetCube(img, cc, size, x, y, black)
		}

	}

	pixels := imagehandling.PNGDataToPixelSlice(img, imgWidth+1, imgHeight+1)

	imagehandling.WritePNGFromPixels(filepath, pixels)

}

func drawOutlines(img *image.RGBA, height, width int, outline color.RGBA) {
	// face 0

	draw.Draw(img, image.Rect(0, height, width, height*2), &image.Uniform{color.RGBA{240, 240, 240, 255}}, image.Point{}, draw.Src)

	// faces 2 and 3
	draw.Draw(img, image.Rect(width*2, height, width*4, height*2), &image.Uniform{color.RGBA{240, 240, 240, 255}}, image.Point{}, draw.Src)

	// line between faces 2 and 3
	d.StraightLine2(width*3, height, width*3, height*2, img, color.RGBA{0, 0, 0, 255})

	// face 4
	draw.Draw(img, image.Rect(width, 0, width*2, height), &image.Uniform{color.RGBA{240, 240, 240, 255}}, image.Point{}, draw.Src)

	// face 5
	draw.Draw(img, image.Rect(width, height*2, width*2, height*3), &image.Uniform{color.RGBA{240, 240, 240, 255}}, image.Point{}, draw.Src)

	// face 1
	draw.Draw(img, image.Rect(width, height, width*2, height*2), &image.Uniform{color.RGBA{240, 240, 240, 255}}, image.Point{}, draw.Src)

}

func noInsetCube(img *image.RGBA, c *cell.CubeCell, cellSize, x, y int, wall color.RGBA) {
	x1, y1 := x, y
	x2 := x1 + cellSize
	y2 := y1 + cellSize

	if c.North == nil {

		d.StraightLine2(x1, y1, x2, y1, img, wall)
	}

	if c.West == nil {

		d.StraightLine2(x1, y1, x1, y2, img, wall)
	}

	if !c.IsLinked(c.East) {

		d.StraightLine2(x2, y1, x2, y2, img, wall)
	}
	if !c.IsLinked(c.South) {

		d.StraightLine2(x1, y2, x2, y2, img, wall)
	}

}

func withInsetCube(img *image.RGBA, cc *cell.CubeCell, cellSize, x, y, inset int, wall color.RGBA) {
	x1, x2, x3, x4, y1, y2, y3, y4 := cellCoordsWithInset(x, y, cellSize, inset)

	if cc.IsLinked(cc.North) {
		d.StraightLine2(x2, y1, x2, y2, img, wall)
		d.StraightLine2(x3, y1, x3, y2, img, wall)
	} else {
		d.StraightLine2(x2, y2, x3, y2, img, wall)

	}

	if cc.IsLinked(cc.South) {
		d.StraightLine2(x2, y3, x2, y4, img, wall)
		d.StraightLine2(x3, y3, x3, y4, img, wall)
	} else {
		d.StraightLine2(x2, y3, x3, y3, img, wall)

	}

	if cc.IsLinked(cc.West) {
		d.StraightLine2(x1, y2, x2, y2, img, wall)
		d.StraightLine2(x1, y3, x2, y3, img, wall)
	} else {
		d.StraightLine2(x2, y2, x2, y3, img, wall)

	}

	if cc.IsLinked(cc.East) {
		d.StraightLine2(x3, y2, x4, y2, img, wall)
		d.StraightLine2(x3, y3, x4, y3, img, wall)
	} else {
		d.StraightLine2(x3, y2, x3, y3, img, wall)

	}
}
