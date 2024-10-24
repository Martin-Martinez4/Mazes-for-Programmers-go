package main

type PolarCell struct {
	row     int
	column  int
	cw      *PolarCell
	ccw     *PolarCell
	inward  *PolarCell
	outward *PolarCell
	links   map[*PolarCell]bool
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

func (pc *PolarCell) Row() int {
	return pc.row
}

func (pc *PolarCell) Column() int {
	return pc.column
}

func (pc *PolarCell) Distances() *Distances {
	distances := CreateDistances(pc)

	// may have to change
	frontier := CreateQueue(40)

	currentCell := pc
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

	if pc.cw != nil {
		neighs = append(neighs, pc.cw)
	}
	if pc.ccw != nil {
		neighs = append(neighs, pc.ccw)
	}
	if pc.inward != nil {
		neighs = append(neighs, pc.inward)
	}
	if pc.outward != nil {
		neighs = append(neighs, pc.outward)
	}

	return neighs
}
