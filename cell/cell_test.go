package cell

import (
	"fmt"
	"testing"
)

func TestLinkCells(t *testing.T) {

	cell1 := CreateBaseCell(0, 0)
	cell2 := CreateBaseCell(1, 0)
	cell3 := CreateBaseCell(2, 0)
	cell4 := CreateBaseCell(3, 0)
	cell5 := CreateBaseCell(0, 1)
	cell6 := CreateBaseCell(0, 2)
	cell7 := CreateBaseCell(0, 3)

	cells := []Cell{cell1, cell2, cell3, cell4, cell5, cell6, cell7}

	for i := 1; i < len(cells); i++ {
		cells[0].Link(cells[i])
	}

	fmt.Println(len(cells[0].Links()))
	if len(cells[0].Links()) != len(cells)-1 {
		t.Errorf("First cell does not have the correct number of links; got: %d expected: %d", len(cells[0].Links()), len(cells)-1)
	}

	c1Links := cell1.Links()
	c2Found := false
	for i := 0; i < len(c1Links); i++ {
		if cell2 == c1Links[i] {
			c2Found = true
		}
	}

	if !c2Found {
		t.Error("cell2 not found in c1links")
	}

	c5Found := false
	for i := 0; i < len(c1Links); i++ {
		if cell5 == c1Links[i] {
			c5Found = true
		}
	}

	if !c5Found {
		t.Error("cell5 not found in c1links")
	}

	cell1.Unlink(cell5)

	c5Found = false
	c1Links = cell1.Links()
	for i := 0; i < len(c1Links); i++ {
		if cell5 == c1Links[i] {
			c5Found = true
		}
	}

	if c5Found {
		t.Error("cell5 found in c1links")
	}

	if len(cell5.links) != 0 {
		t.Errorf("expected length of cell5 links to be 0 got: %d", len(cell5.links))
	}

}
