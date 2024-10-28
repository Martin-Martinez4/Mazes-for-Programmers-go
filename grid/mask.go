package grid

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/imagehandling"
)

type Mask struct {
	rows     int
	columns  int
	grid     [][]bool
	numberOn int
}

func CreateMask(rows, columns int) *Mask {
	mask := &Mask{rows: rows, columns: columns}
	mask.AllOnGrid(rows, columns)

	return mask
}

func CreateMaskFromFile(filepath string) {

	MaskFromFile(filepath)
}

func (m *Mask) AllOnGrid(rows, columns int) {
	grid := make([][]bool, rows)

	for row := 0; row < rows; row++ {
		grid[row] = make([]bool, columns)

		for column := 0; column < columns; column++ {

			grid[row][column] = true
		}
	}

	m.numberOn = rows * columns
	m.grid = grid
}

func MaskFromFile(filepath string) *Mask {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	line, _, err := fileReader.ReadLine()
	if err != nil {
		panic("File Error")
	}

	rows, columns := getRowsColumnsFromLine(line)

	mask := CreateMask(rows, columns)

	for row := 0; row < rows; row++ {
		line, _, err := fileReader.ReadLine()
		if len(line) < columns {
			panic(fmt.Sprintf("line %d in file %s is too short", row, filepath))
		}
		if err != nil {
			break
		}

		for column := 0; column < columns; column++ {
			if line[column] == 'X' || line[column] == 'x' {
				mask.turnOff(row, column)
			}
		}
	}

	return mask
}

func getRowsColumnsFromLine(line []byte) (rows int, columns int) {
	r, err := regexp.Compile("[0-9]+[X|x][0-9]+")
	if err != nil {
		panic(err)
	}

	if !r.MatchString(string(line)) {
		panic("Malformed mask file, dimensions must be first line")
	}

	dimensions := strings.Split(strings.ToLower(string(line)), "x")
	rs, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic("Problem getting rows")
	}

	cs, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic("Problem getting columns")
	}

	return rs, cs
}

func MaskFromPNG(filepath string) *Mask {
	img, w, h := imagehandling.ReadPNG(filepath)
	pixels := imagehandling.PNGDataToPixelSlice(img, w, h)

	m := CreateMask(h, w)

	for row := 0; row < h; row++ {
		for column := 0; column < w; column++ {
			p := pixels[row][column]
			if p.Equal(&imagehandling.Pixel{R: 0, G: 0, B: 0, A: 255}) {
				m.turnOff(row, column)
			}
		}
	}

	return m
}

func (m *Mask) turnOff(row, column int) {
	if m.grid[row][column] {

		m.grid[row][column] = false
		m.numberOn -= 1
	}
}

func (m *Mask) turnOn(row, column int) {
	if !m.grid[row][column] {

		m.grid[row][column] = true
		m.numberOn += 1
	}
}

func (m *Mask) isOn(row, column int) bool {
	return m.grid[row][column]
}

func (m *Mask) countOn() int {
	return m.numberOn
}

func (m *Mask) randomLocation() (row int, column int) {

	randRow := rand.Intn(len(m.grid))
	randColumn := rand.Intn(len(m.grid[0]))

	for !m.grid[randRow][randColumn] {

		randRow = rand.Intn(len(m.grid))
		randColumn = rand.Intn(len(m.grid[0]))
	}

	return row, column
}
