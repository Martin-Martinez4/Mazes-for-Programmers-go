package cell

type Cell3D struct {
	*PrimeCell
	North *Cell3D
	South *Cell3D
	East  *Cell3D
	West  *Cell3D
	links map[*Cell3D]bool
	level int
	Up    *Cell3D
	Down  *Cell3D
}

func CreateCell3D(level, row, column int) *Cell3D {
	return &Cell3D{
		PrimeCell: &PrimeCell{
			row:    row,
			column: column,
			weight: 1,
		},
		level: level,
		links: map[*Cell3D]bool{},
	}
}

func (c3D *Cell3D) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range c3D.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (c3D *Cell3D) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*Cell3D)
	if !ok {
		return
	}

	_, ok = c3D.links[c2]
	if !ok {
		c3D.links[c2] = true
	}

	_, ok = c2.links[c3D]
	if !ok {

		c2.links[c3D] = true
	}

}

func (c3D *Cell3D) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*Cell3D)
	if !ok {
		return
	}

	delete(c3D.links, c2)
	delete(c2.links, c3D)
}

func (c3D *Cell3D) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*Cell3D)
	if !ok {
		return false
	}
	_, ok = c3D.links[c2]

	return ok
}

func (c3D *Cell3D) Distances() *Distances {
	distances := CreateDistances(c3D)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := c3D
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

			c2, ok := frontier.Pop().(*Cell3D)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*Cell3D)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}

func (c3D *Cell3D) Neighbors() []Cell {
	neighs := []Cell{}

	if c3D.North != nil {
		neighs = append(neighs, c3D.North)
	}
	if c3D.South != nil {
		neighs = append(neighs, c3D.South)
	}
	if c3D.West != nil {
		neighs = append(neighs, c3D.West)
	}
	if c3D.East != nil {
		neighs = append(neighs, c3D.East)
	}
	if c3D.Up != nil {
		neighs = append(neighs, c3D.Up)
	}

	if c3D.Down != nil {
		neighs = append(neighs, c3D.Down)
	}

	return neighs
}
