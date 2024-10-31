package grid

import (
	"testing"
)

func TestGrid3D(t *testing.T) {
	g3d := CreateGrid3D(2, 4, 4)

	for lv := 0; lv < 2; lv++ {
		for r := 0; r < 4; r++ {
			for c := 0; c < 4; c++ {
				c := g3d.GetCell(lv, r, c)
				if c == nil {
					t.Error("Cell was nil, when it should not have been")
				}

				if len(c.Neighbors()) == 0 {
					t.Errorf("Every Cell should have neighbors; cell (%d, %d, %d) does not", lv, r, c)
				}

				neighs := c.Neighbors()
				for i := 0; i < len(neighs); i++ {
					if neighs[i] == nil {
						t.Error("A neighbor was nil when nil should never be a neighbor")
					}
				}

				c.Links()
			}

		}
	}

	if g3d.GetCell(0, 0, 0) != g3d.Grid[0][0][0] {
		t.Error("First cell did not match")
	}

	neighs := g3d.GetCell(0, 0, 0).Neighbors()
	if len(neighs) != 3 {
		t.Error("wrong number of nieghbors for first cell")
	}

	if g3d.GetCell(100, 100, 100) != nil {
		t.Error("Cell was not nil, when it should have been")
	}

}
