package imagehandling

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func WritePNG(filepath string, width, height int) {

	black := &Pixel{0, 0, 0, 255}

	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{R: uint8(black.R), G: uint8(black.G), B: uint8(black.B), A: uint8(black.A)})
		}
	}

	f, err := os.Create(filepath + ".png")
	if err != nil {
		panic("Failed to create png file")
	}
	png.Encode(f, img)
}

func WritePNGFromPixels(filepath string, pixels [][]*Pixel) {

	width := len(pixels[0])
	height := len(pixels)

	upperLeft := image.Point{0, 0}
	lowerRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upperLeft, lowerRight})

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			c := pixels[x][y]

			img.Set(y, x, color.RGBA{R: uint8(c.R), G: uint8(c.G), B: uint8(c.B), A: uint8(c.A)})
		}
	}

	f, err := os.Create(filepath + ".png")
	if err != nil {
		panic("Failed to create png file")
	}
	png.Encode(f, img)

}
