package cell

type CubeCell struct {
	*PrimeCell
	face  int
	North *CubeCell
	South *CubeCell
	East  *CubeCell
	West  *CubeCell
	links map[*CubeCell]bool
}

func CreateCubeCell(face, row, column int) *CubeCell {
	return &CubeCell{
		PrimeCell: CreatePrimeCell(row, column),
		face:      face,
		links:     map[*CubeCell]bool{},
	}
}

func (cc *CubeCell) Face() int {
	return cc.face
}

func (c *CubeCell) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*CubeCell)
	if !ok {
		return
	}

	_, ok = c.links[c2]
	if !ok {
		c.links[c2] = true
	}

	_, ok = c2.links[c]
	if !ok {

		c2.links[c] = true
	}

}

func (c *CubeCell) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*CubeCell)
	if !ok {
		return
	}

	delete(c.links, c2)
	delete(c2.links, c)
}

func (c *CubeCell) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range c.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (c *CubeCell) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*CubeCell)
	if !ok {
		return false
	}
	_, ok = c.links[c2]

	return ok
}

func (c *CubeCell) Neighbors() []Cell {

	var cells []Cell

	if c.North != nil {
		cells = append(cells, c.North)
	}
	if c.South != nil {
		cells = append(cells, c.South)
	}
	if c.East != nil {
		cells = append(cells, c.East)
	}
	if c.West != nil {
		cells = append(cells, c.West)
	}

	return cells
}

func (c *CubeCell) Distances() *Distances {
	distances := CreateDistances(c)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := c
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

			c2, ok := frontier.Pop().(*CubeCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*CubeCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}
