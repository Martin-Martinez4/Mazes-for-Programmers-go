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
