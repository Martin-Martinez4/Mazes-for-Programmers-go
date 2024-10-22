package imagehandling

import (
	"fmt"
	"testing"
)

func TestReadPNG(t *testing.T) {

	// path has to be relative to the path and not go.mod file
	tests := []struct {
		name     string
		filepath string
		width    int
		height   int
	}{
		{
			name:     "10x10 square border all around",
			filepath: "../images/test/do_not_delete.png",
			width:    10,
			height:   10,
		},
		{
			name:     "20x10 square border all around and in middle",
			filepath: "../images/test/do_not_delete_2.png",
			width:    20,
			height:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, width, height := ReadPNG(tt.filepath)

			if width != tt.width {
				t.Errorf("Width does not match; got: %d, expected: %d", width, tt.width)
			}

			if height != tt.height {
				t.Errorf("height does not match; got: %d, expected: %d", height, tt.height)
			}
		})
	}
}

func TestPNGDataToPixelSlice(t *testing.T) {

	black := &Pixel{0, 0, 0, 255}
	white := &Pixel{255, 255, 255, 255}

	// path has to be relative to the path and not go.mod file
	tests := []struct {
		name     string
		filepath string
		pixels   [][]*Pixel
	}{
		{
			name:     "10x10 square border all around",
			filepath: "../images/test/do_not_delete.png",
			pixels: [][]*Pixel{
				{black, black, black, black, black, black, black, black, black, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, black, black, black, black, black, black, black, black, black},
			},
		},
		{
			name:     "20x10 square border all around and in middle",
			filepath: "../images/test/do_not_delete_2.png",
			pixels: [][]*Pixel{
				{black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, white, black, white, white, white, white, white, white, white, white, black},
				{black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black, black},
			},
		},
		{
			name:     "10x10 square border all around except top",
			filepath: "../images/test/do_not_delete_3.png",
			pixels: [][]*Pixel{
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, white, white, white, white, white, white, white, white, black},
				{black, black, black, black, black, black, black, black, black, black},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			img, width, height := ReadPNG(tt.filepath)
			pixels := PNGDataToPixelSlice(img, width, height)

			if len(pixels) != len(tt.pixels) {
				t.Errorf("height did not match; got: %d, %d wanted: %d", len(pixels), len(tt.pixels), height)
			}

			if len(pixels[0]) != len(tt.pixels[0]) {
				t.Errorf("width did not match; got: %d, %d wanted: %d", len(pixels[0]), len(tt.pixels[0]), width)
			}

			for h := 0; h < height; h++ {
				for w := 0; w < width; w++ {
					if !tt.pixels[h][w].Equal(pixels[h][w]) {
						fmt.Println(pixels[h][w])
						t.Errorf("Pixel colors did not match at pixel: (%d,%d)", w, h)
					}
				}
			}
		})
	}
}
