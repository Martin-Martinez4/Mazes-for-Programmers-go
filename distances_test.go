package main

import (
	"testing"
)

func TestDistances(t *testing.T) {
	grid1 := CreateShape(2, 2)

	g1c00 := grid1.grid[0][0]
	distances := CreateDistances(g1c00)

	if distances.cells[g1c00] != 0 {
		t.Error("Expected cell 0, 0 to have distance of 0")
	}

	grid1 = CreateShape(2, 2)

	g1c00 = grid1.grid[0][0]
	g1c01 := grid1.grid[0][1]
	g1c10 := grid1.grid[1][0]
	g1c11 := grid1.grid[1][1]

	g1c00.Link(g1c01)
	g1c00.Link(g1c10)

	g1c01.Link(g1c11)

	distances = g1c00.Distances()

	// if distances.cells[g1c00] != 0 {
	// 	t.Error("Expected cell 0, 0 to have distance of 0")
	// }

	// if distances.cells[g1c01] != 1 {
	// 	t.Errorf("Expected cell 0, 1 to have distance of 0; got: %d", distances.cells[g1c01])
	// }

	// if distances.cells[g1c11] != 2 {
	// 	t.Errorf("Expected cell 1, 1 to have distance of 2; got: %d", distances.cells[g1c01])
	// }

	grid1 = CreateShape(3, 3)

	g1c00 = grid1.grid[0][0]
	g1c01 = grid1.grid[0][1]
	g1c02 := grid1.grid[0][2]

	g1c10 = grid1.grid[1][0]
	g1c11 = grid1.grid[1][1]
	g1c12 := grid1.grid[1][2]

	g1c20 := grid1.grid[2][0]
	g1c21 := grid1.grid[2][1]
	g1c22 := grid1.grid[2][2]

	type test struct {
		name     string
		cell     Cell
		distance int
	}

	tests := []test{}

	currentLevel := 0
	// 0
	tests = append(tests, test{"g1c00", g1c00, currentLevel})
	// 1
	g1c00.Link(g1c01)
	g1c00.Link(g1c02)
	currentLevel++
	tests = append(tests, test{"g1c01", g1c01, currentLevel})
	tests = append(tests, test{"g1c02", g1c02, currentLevel})

	// 2
	g1c01.Link(g1c10)
	g1c01.Link(g1c11)
	currentLevel++
	tests = append(tests, test{"g1c10", g1c10, currentLevel})
	tests = append(tests, test{"g1c11", g1c11, currentLevel})

	// 3
	g1c11.Link(g1c12)
	currentLevel++
	tests = append(tests, test{"g1c12", g1c12, currentLevel})

	// 4
	g1c12.Link(g1c20)
	currentLevel++
	tests = append(tests, test{"g1c20", g1c20, currentLevel})

	// 5
	g1c20.Link(g1c21)
	currentLevel++
	tests = append(tests, test{"g1c21", g1c21, currentLevel})

	// 6
	g1c21.Link(g1c22)
	currentLevel++
	tests = append(tests, test{"g1c22", g1c22, currentLevel})

	distances = g1c00.Distances()

	for i := 0; i < len(tests); i++ {
		cell := tests[i].cell
		dist := tests[i].distance

		if distances.cells[cell] != dist {
			t.Errorf("%s failed: distance expected %d; got %d", tests[i].name, dist, distances.cells[cell])
		}
	}

}
