package draw

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

func StraightLine(x1, y1, x2, y2 int, pixels [][]*imagehandling.Pixel, color imagehandling.Pixel) {

	var xmin, xmax, ymin, ymax int

	if x1 < x2 {
		xmin, xmax = x1, x2
	} else {
		xmin, xmax = x2, x1
	}

	if y1 < y2 {
		ymin, ymax = y1, y2
	} else {
		ymin, ymax = y2, y1
	}

	slope := float64(y2-y1) / float64(x2-x1)
	yInter := float64(y1) - (slope * float64(x1))

	if float64(x2-x1) == 0 {
		if float64(y2-y1) == 0 {
			// 0/0, nothing happens
			return
		} else {
			// slope is 0; veritcal line
			for y := ymin; y <= ymax; y++ {
				pixels[xmin][y] = &color
			}
		}
	} else {

	}

}
