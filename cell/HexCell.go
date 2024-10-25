package cell

type HexCell struct {
	*PrimeCell
	North     *HexCell
	South     *HexCell
	NorthEast *HexCell
	NorthWest *HexCell
	SouthWest *HexCell
	SouthEast *HexCell
	links     map[*HexCell]bool
}

func CreateHexCell(row, column int) *HexCell {
	return &HexCell{
		PrimeCell: &PrimeCell{
			row:    row,
			column: column,
		},
		links: map[*HexCell]bool{},
	}
}

func (hc *HexCell) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range hc.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (hc *HexCell) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*HexCell)
	if !ok {
		return
	}

	_, ok = hc.links[c2]
	if !ok {
		hc.links[c2] = true
	}

	_, ok = c2.links[hc]
	if !ok {

		c2.links[hc] = true
	}

}

func (hc *HexCell) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*HexCell)
	if !ok {
		return
	}

	delete(hc.links, c2)
	delete(c2.links, hc)
}

func (hc *HexCell) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*HexCell)
	if !ok {
		return false
	}
	_, ok = hc.links[c2]

	return ok
}

func (hc *HexCell) Distances() *Distances {
	distances := CreateDistances(hc)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := hc
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

			c2, ok := frontier.Pop().(*HexCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*HexCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}

func (hc *HexCell) Neighbors() []Cell {
	neighs := []Cell{}

	if hc.North != nil {
		neighs = append(neighs, hc.North)
	}
	if hc.South != nil {
		neighs = append(neighs, hc.South)
	}
	if hc.NorthWest != nil {
		neighs = append(neighs, hc.NorthWest)
	}
	if hc.NorthEast != nil {
		neighs = append(neighs, hc.NorthEast)
	}
	if hc.SouthWest != nil {
		neighs = append(neighs, hc.SouthWest)
	}
	if hc.SouthEast != nil {
		neighs = append(neighs, hc.SouthEast)
	}

	return neighs
}
