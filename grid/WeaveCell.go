package grid

import "github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"

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
		ov.Link(uc.North)
		ov.Link(uc.South)

	}

	return uc
}

func (uc *OverUnderCell) HorizontalPassage() bool {
	return uc.East != nil || uc.West != nil
}

func (uc *OverUnderCell) VerticalPassage() bool {
	return uc.North != nil || uc.South != nil
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

func (oc *OverUnderCell) HorizontalPassage() bool {
	return oc.IsLinked(oc.East) && oc.IsLinked(oc.West) && !oc.IsLinked(oc.East) && !oc.IsLinked(oc.West)
}

func (oc *OverUnderCell) VerticalPassage() bool {
	return !oc.IsLinked(oc.East) && !oc.IsLinked(oc.West) && oc.IsLinked(oc.East) && oc.IsLinked(oc.West)
}

func (oc *OverUnderCell) Link(cel cell.Cell) {
	var neighbor *cell.BaseCell
	c, ok := cel.(*cell.BaseCell)
	if !ok {
		return
	}

	if oc.North != nil && oc.North == c.South {
		neighbor = oc.North
	}

	if oc.South != nil && oc.South == c.North {
		neighbor = oc.South
	}

	if oc.East != nil && oc.East == c.West {
		neighbor = oc.East
	}
	if oc.West != nil && oc.West == c.East {
		neighbor = oc.West
	}

	c2, ok := cel.(*OverUnderCell)
	if !ok {
		return
	}

	if neighbor != nil {
		oc.grid.TunnelUnder(c2)
	}
}
