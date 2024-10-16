package main

import (
	"testing"
)

func TestLinkCells(t *testing.T) {

	cell1 := CreateCell(0, 0)
	cell2 := CreateCell(1, 0)
	cell3 := CreateCell(2, 0)
	cell4 := CreateCell(3, 0)
	cell5 := CreateCell(0, 1)
	cell6 := CreateCell(0, 2)
	cell7 := CreateCell(0, 3)

	cells := []*Cell{cell1, cell2, cell3, cell4, cell5, cell6, cell7}

	for i := 1; i < len(cells); i++ {
		cells[0].link(cells[i])
	}

	if len(cells[0].links) != len(cells)-1 {
		t.Errorf("First cell does not have the correct number of links; got: %d expected: %d", len(cells[0].links), len(cells)-1)
	}

	c1Links := cell1.getLinks()
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

	cell1.unlink(cell5)

	c5Found = false
	c1Links = cell1.getLinks()
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
