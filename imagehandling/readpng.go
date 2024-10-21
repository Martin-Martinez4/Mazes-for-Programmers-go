package imagehandling

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func (p *Pixel) equal(p2 *Pixel) bool {
	return p.R == p2.R && p.G == p2.G && p.B == p2.B
}

func ReadPNG(filepath string) (img image.Image, width, height int) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	i, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: File could not be decoded")
		os.Exit(1)
	}

	bounds := i.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	return i, w, h
}

func PNGDataToPixelSlice(img image.Image, width, height int) [][]*Pixel {

	pixels := make([][]*Pixel, height)

	for y := 0; y < height; y++ {
		pixels[y] = make([]*Pixel, width)

		for x := 0; x < width; x++ {
			// values are never more than 255, so casting to int should be safe
			r, g, b, a := img.At(x, y).RGBA()
			pixels[y][x] = &Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257), A: int(a / 257)}

		}
	}
	return pixels
}
