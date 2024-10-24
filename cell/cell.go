package cell

type Cell interface {
	Links() []Cell
	Link(Cell)
	Unlink(Cell)
	IsLinked(Cell) bool

	Row() int
	Column() int

	Neighbors() []Cell
	Distances() *Distances
}

type BaseCell struct {
	*PrimeCell
	North *BaseCell
	South *BaseCell
	East  *BaseCell
	West  *BaseCell
	links map[*BaseCell]bool
}

func CreateBaseCell(row, column int) *BaseCell {

	return &BaseCell{
		PrimeCell: &PrimeCell{
			row:    row,
			column: column,
		},
		links: map[*BaseCell]bool{},
	}

}

func (c *BaseCell) Link(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*BaseCell)
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

func (c *BaseCell) Unlink(cell Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*BaseCell)
	if !ok {
		return
	}

	delete(c.links, c2)
	delete(c2.links, c)
}

func (c *BaseCell) Links() []Cell {
	cellLinks := []Cell{}

	for key, _ := range c.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (c *BaseCell) IsLinked(cell Cell) bool {

	if cell == nil {
		return false
	}
	c2, ok := cell.(*BaseCell)
	if !ok {
		return false
	}
	_, ok = c.links[c2]

	return ok
}

func (c *BaseCell) Neighbors() []Cell {

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

func (c *BaseCell) Distances() *Distances {
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

			c2, ok := frontier.Pop().(*BaseCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*BaseCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}
