package cell

type TriangleCell struct {
	*PrimeCell
	Upright bool
	North   *TriangleCell
	South   *TriangleCell
	East    *TriangleCell
	West    *TriangleCell
	links   map[*TriangleCell]bool
}

func CreateTriangleCell(row, column int) *TriangleCell {

	upright := false
	if (row+column)%2 == 0 {
		upright = true
	}
	return &TriangleCell{
		PrimeCell: &PrimeCell{
			row:    row,
			column: column,
			weight: 1,
		},
		Upright: upright,
		links:   map[*TriangleCell]bool{},
	}
}

func (tc *TriangleCell) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range tc.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (tc *TriangleCell) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*TriangleCell)
	if !ok {
		return
	}

	_, ok = tc.links[c2]
	if !ok {
		tc.links[c2] = true
	}

	_, ok = c2.links[tc]
	if !ok {

		c2.links[tc] = true
	}

}

func (tc *TriangleCell) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*TriangleCell)
	if !ok {
		return
	}

	delete(tc.links, c2)
	delete(c2.links, tc)
}

func (tc *TriangleCell) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*TriangleCell)
	if !ok {
		return false
	}
	_, ok = tc.links[c2]

	return ok
}

func (tc *TriangleCell) Distances() *Distances {
	distances := CreateDistances(tc)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := tc
	distances.Cells[currentCell] = 0

	for currentCell != nil {
		// may have to change
		newFrontier := CreateQueue(40)

		for currentCell != nil {

			for key, _ := range currentCell.links {

				_, ok := distances.Cells[key]
				if !ok {
					newFrontier.Push(key)
					distances.Cells[key] = distances.Cells[currentCell] + 1
				}

			}

			c2, ok := frontier.Pop().(*TriangleCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*TriangleCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}

func (tc *TriangleCell) Neighbors() []Cell {
	neighs := []Cell{}

	if (tc.North != nil) && !tc.Upright {
		neighs = append(neighs, tc.North)
	}
	if (tc.South != nil) && tc.Upright {
		neighs = append(neighs, tc.South)
	}
	if tc.West != nil {
		neighs = append(neighs, tc.West)
	}
	if tc.East != nil {
		neighs = append(neighs, tc.East)
	}

	return neighs
}
