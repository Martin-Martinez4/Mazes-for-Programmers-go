package grid

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
)

type OverUnderCell struct {
	*cell.PrimeCell
	North   *OverUnderCell
	South   *OverUnderCell
	East    *OverUnderCell
	West    *OverUnderCell
	links   map[*OverUnderCell]bool
	grid    *WeaveGrid
	IsUnder bool
}

func CreateOverCell(row, column int, grid *WeaveGrid) *OverUnderCell {
	return &OverUnderCell{
		PrimeCell: cell.CreatePrimeCell(row, column),
		links:     map[*OverUnderCell]bool{},
		grid:      grid,
		IsUnder:   false,
	}

}

func CreateUnderCell(ov *OverUnderCell) *OverUnderCell {
	uc := &OverUnderCell{
		PrimeCell: cell.CreatePrimeCell(ov.Row(), ov.Column()),
		links:     map[*OverUnderCell]bool{},
		IsUnder:   true,
	}

	if ov.HorizontalPassage() {
		uc.North = ov.North
		ov.North.South = uc
		uc.South = ov.South
		ov.South.North = uc
		uc.Link(ov.North)
		uc.Link(ov.South)

	} else {
		uc.East = ov.East
		ov.East.West = uc
		uc.West = ov.West
		ov.West.East = uc

		uc.Link(ov.East)
		uc.Link(ov.West)
	}

	return uc
}

func (oc *OverUnderCell) HorizontalPassage() bool {
	if oc.IsUnder {

		return oc.East != nil || oc.West != nil
	} else {
		return oc.IsLinked(oc.East) && oc.IsLinked(oc.West) && !oc.IsLinked(oc.North) && !oc.IsLinked(oc.South)

	}
}

func (oc *OverUnderCell) VerticalPassage() bool {
	if oc.IsUnder {

		return oc.North != nil || oc.South != nil
	} else {
		return !oc.IsLinked(oc.East) && !oc.IsLinked(oc.West) && oc.IsLinked(oc.North) && oc.IsLinked(oc.South)

	}
}

func (oc *OverUnderCell) CanTunnelNorth() bool {
	if oc.North != nil {

		return oc.North.North != nil && oc.HorizontalPassage()
	}

	return false
}

func (oc *OverUnderCell) CanTunnelSouth() bool {
	if oc.South != nil {

		return oc.South.South != nil && oc.HorizontalPassage()
	}

	return false
}

func (oc *OverUnderCell) CanTunnelEast() bool {
	if oc.East != nil {

		return oc.East.East != nil && oc.VerticalPassage()
	}

	return false

}

func (oc *OverUnderCell) CanTunnelWest() bool {
	if oc.West != nil {

		return oc.West.West != nil && oc.VerticalPassage()
	}

	return false

}

func (oc *OverUnderCell) Link(cel cell.Cell) {
	var neighbor *OverUnderCell

	c, ok := cel.(*OverUnderCell)
	if !ok {
		panic("Failed to cast to OverUnderCell")
	}

	if oc.North != nil && oc.North == c.South {
		neighbor = oc.North
	} else if oc.South != nil && oc.South == c.North {
		neighbor = oc.South
	} else if oc.East != nil && oc.East == c.West {
		neighbor = oc.East
	} else if oc.West != nil && oc.West == c.East {
		neighbor = oc.West
	}

	if neighbor != nil {
		oc.grid.TunnelUnder(neighbor)
	} else {
		if cel == nil {
			return
		}
		c2, ok := cel.(*OverUnderCell)
		if !ok {
			panic("Failed to cast to OverUnderCell")
		}

		_, ok = oc.links[c2]
		if !ok {
			oc.links[c2] = true
		}

		_, ok = c2.links[oc]
		if !ok {

			c2.links[oc] = true
		}
	}
}

func (oc *OverUnderCell) Links() []cell.Cell {
	cellLinks := []cell.Cell{}

	for key, _ := range oc.links {
		cellLinks = append(cellLinks, key)
	}

	return cellLinks
}

func (pc *OverUnderCell) Unlink(cell cell.Cell) {
	if cell == nil {
		return
	}
	c2, ok := cell.(*OverUnderCell)
	if !ok {
		return
	}

	delete(pc.links, c2)
	delete(c2.links, pc)
}

func (pc *OverUnderCell) IsLinked(c cell.Cell) bool {

	if c == nil {
		return false
	}
	c2, ok := c.(*OverUnderCell)
	if !ok {
		return false
	}
	_, ok = pc.links[c2]

	return ok
}

func (pc *OverUnderCell) Distances() *cell.Distances {
	distances := cell.CreateDistances(pc)

	// may have to change
	frontier := cell.CreateQueue(40)

	currentCell := pc
	distances.Cells[currentCell] = 0

	for currentCell != nil {
		// may have to change
		newFrontier := cell.CreateQueue(40)

		for currentCell != nil {

			for key, _ := range currentCell.links {

				_, ok := distances.Cells[key]
				if !ok {
					newFrontier.Push(key)
					distances.Cells[key] = distances.Cells[currentCell] + 1
				}

			}

			c2, ok := frontier.Pop().(*OverUnderCell)
			if !ok && c2 != nil {
				return distances
			}
			currentCell = c2

		}

		frontier = newFrontier
		c2, ok := frontier.Pop().(*OverUnderCell)
		if !ok && c2 != nil {
			return distances
		}
		currentCell = c2

	}

	return distances
}

func (oc *OverUnderCell) Neighbors() []cell.Cell {
	var cells []cell.Cell

	if oc.North != nil {
		cells = append(cells, oc.North)
	}
	if oc.South != nil {
		cells = append(cells, oc.South)
	}
	if oc.East != nil {
		cells = append(cells, oc.East)
	}
	if oc.West != nil {
		cells = append(cells, oc.West)
	}

	if oc.CanTunnelNorth() {
		cells = append(cells, oc.North.North)
	}

	if oc.CanTunnelSouth() {
		cells = append(cells, oc.South.South)
	}
	if oc.CanTunnelEast() {
		cells = append(cells, oc.East.East)
	}
	if oc.CanTunnelWest() {
		cells = append(cells, oc.West.West)
	}

	return cells

}
