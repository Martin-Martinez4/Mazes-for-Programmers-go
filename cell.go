package main

type Cell struct {
	row    int
	column int
	north  *Cell
	south  *Cell
	east   *Cell
	west   *Cell
	links  map[*Cell]bool
}

func CreateCell(row, column int) *Cell {
	return &Cell{
		row:    row,
		column: column,
		links:  map[*Cell]bool{},
	}

}

func (c *Cell) link(cell *Cell) {
	if cell == nil {
		return
	}

	_, ok := c.links[cell]
	if !ok {
		c.links[cell] = true
	}

	_, ok = cell.links[c]
	if !ok {

		cell.links[c] = true
	}

}

func (c *Cell) unlink(cell *Cell) {
	if cell == nil {
		return
	}

	delete(c.links, cell)
	delete(cell.links, c)
}

func (c *Cell) getLinks() []*Cell {
	cellLinks := make([]*Cell, len(c.links))

	for key, _ := range c.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (c *Cell) isLinked(cell *Cell) bool {
	_, ok := c.links[cell]

	return ok
}

func (c *Cell) neighbors() []*Cell {

	var cells []*Cell

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

func (c *Cell) distances() *Distances {
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

			currentCell = frontier.Pop()

		}

		frontier = newFrontier
		currentCell = frontier.Pop()

	}

	return distances
}
