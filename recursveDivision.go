package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

func RecursveDivision(sh grid.ShapeHolder) {

	// for row := 0; row < len(g); row++ {
	// 	for column := 0; column < len(g[row]); column++ {
	for c := range sh.EachCell() {

		neighs := c.Neighbors()

		for n := 0; n < len(neighs); n++ {
			// may need to add option to not make it bidirectional
			c.Link(neighs[n])
		}
	}

	// 	}
	// }

	divide(sh, 0, 0, sh.Rows(), sh.Columns())
}

func RecursveDivisionRooms(sh grid.ShapeHolder) {

	// for row := 0; row < len(g); row++ {
	// 	for column := 0; column < len(g[row]); column++ {
	for c := range sh.EachCell() {

		neighs := c.Neighbors()

		for n := 0; n < len(neighs); n++ {
			// may need to add option to not make it bidirectional
			c.Link(neighs[n])
		}
	}

	// 	}
	// }

	divideWithRooms(sh, 0, 0, sh.Rows(), sh.Columns())
}

func divide(sh grid.ShapeHolder, row, column, height, width int) {
	if height <= 1 || width <= 1 || height < 5 && width < 5 && rand.Intn(4) == 0 {
		return
	}

	if height > width {
		divideHorizontally(sh, row, column, height, width)
	} else {
		divideVertically(sh, row, column, height, width)
	}
}

func divideWithRooms(sh grid.ShapeHolder, row, column, height, width int) {
	if height <= 1 || width <= 1 {
		return
	}

	if height > width {
		divideHorizontallyWithRooms(sh, row, column, height, width)
	} else {
		divideVerticallyWithRooms(sh, row, column, height, width)
	}
}

func divideHorizontally(sh grid.ShapeHolder, row, column, height, width int) {
	southDivide := rand.Intn(height - 1)
	passgeAt := rand.Intn(width)

	for w := 0; w < width; w++ {
		if passgeAt != w {
			c := sh.GetShape().GetCell(row+southDivide, column+w).(*cell.BaseCell)
			c.Unlink(c.South)

		}
	}
	divide(sh, row, column, southDivide+1, width)
	divide(sh, row+southDivide+1, column, height-southDivide-1, width)
}
func divideVertically(sh grid.ShapeHolder, row, column, height, width int) {
	eastDivide := rand.Intn(width - 1)
	passgeAt := rand.Intn(height)

	for h := 0; h < height; h++ {
		if passgeAt != h {
			c := sh.GetShape().GetCell(row+h, column+eastDivide).(*cell.BaseCell)
			c.Unlink(c.East)

		}
	}
	divideWithRooms(sh, row, column, height, eastDivide+1)
	divideWithRooms(sh, row, column+eastDivide+1, height, width-eastDivide-1)
}

func divideHorizontallyWithRooms(sh grid.ShapeHolder, row, column, height, width int) {
	southDivide := rand.Intn(height - 1)
	passgeAt := rand.Intn(width)

	for w := 0; w < width; w++ {
		if passgeAt != w {
			c := sh.GetShape().GetCell(row+southDivide, column+w).(*cell.BaseCell)
			c.Unlink(c.South)

		}
	}
	divide(sh, row, column, southDivide+1, width)
	divide(sh, row+southDivide+1, column, height-southDivide-1, width)
}
func divideVerticallyWithRooms(sh grid.ShapeHolder, row, column, height, width int) {
	eastDivide := rand.Intn(width - 1)
	passgeAt := rand.Intn(height)

	for h := 0; h < height; h++ {
		if passgeAt != h {
			c := sh.GetShape().GetCell(row+h, column+eastDivide).(*cell.BaseCell)
			c.Unlink(c.East)

		}
	}
	divideWithRooms(sh, row, column, height, eastDivide+1)
	divideWithRooms(sh, row, column+eastDivide+1, height, width-eastDivide-1)
}
