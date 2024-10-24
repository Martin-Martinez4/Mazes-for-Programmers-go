package cell

type PolarCell struct {
	*PrimeCell
	Cw      *PolarCell
	Ccw     *PolarCell
	Inward  *PolarCell
	Outward []*PolarCell
	links   map[*PolarCell]bool
}

func CreatePolarCell(row, column int) *PolarCell {
	return &PolarCell{
		PrimeCell: &PrimeCell{
			row:    row,
			column: column,
		},
		links: map[*PolarCell]bool{},
	}
}

func (pc *PolarCell) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range pc.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (pc *PolarCell) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*PolarCell)
	if !ok {
		return
	}

	_, ok = pc.links[c2]
	if !ok {
		pc.links[c2] = true
	}

	_, ok = c2.links[pc]
	if !ok {

		c2.links[pc] = true
	}

}

func (pc *PolarCell) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*PolarCell)
	if !ok {
		return
	}

	delete(pc.links, c2)
	delete(c2.links, pc)
}

func (pc *PolarCell) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*PolarCell)
	if !ok {
		return false
	}
	_, ok = pc.links[c2]

	return ok
}

func (pc *PolarCell) Distances() *Distances {
	distances := CreateDistances(pc)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := pc
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

			c2, ok := frontier.Pop().(*PolarCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*PolarCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}

func (pc *PolarCell) Neighbors() []Cell {
	neighs := []Cell{}

	if pc.Cw != nil {
		neighs = append(neighs, pc.Cw)
	}
	if pc.Ccw != nil {
		neighs = append(neighs, pc.Ccw)
	}
	if pc.Inward != nil {
		neighs = append(neighs, pc.Inward)
	}
	if pc.Outward != nil {
		for out := 0; out < len(pc.Outward); out++ {

			neighs = append(neighs, pc.Outward[out])
		}
	}

	return neighs
}
