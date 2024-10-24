package main

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
	row    int
	column int
	north  *BaseCell
	south  *BaseCell
	east   *BaseCell
	west   *BaseCell
	links  map[*BaseCell]bool
}

func CreateBaseCell(row, column int) *BaseCell {
	return &BaseCell{
		row:    row,
		column: column,
		links:  map[*BaseCell]bool{},
	}

}

func (c *BaseCell) Row() int {
	return c.row
}

func (c *BaseCell) Column() int {
	return c.column
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

	if c.north != nil {
		cells = append(cells, c.north)
	}
	if c.south != nil {
		cells = append(cells, c.south)
	}
	if c.east != nil {
		cells = append(cells, c.east)
	}
	if c.west != nil {
		cells = append(cells, c.west)
	}

	// return []*Cell{c.north, c.south, c.east, c.west}
	return cells
}

func (c *BaseCell) Distances() *Distances {
	distances := CreateDistances(c)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := c
	distances.cells[currentCell] = 0

	for currentCell != nil {
		// may have to change
		newFrontier := CreateQueue(40)

		for currentCell != nil {

			for key, _ := range currentCell.links {

				_, ok := distances.cells[key]
				if !ok {
					newFrontier.Push(key)
					distances.cells[key] = distances.cells[currentCell] + 1
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
