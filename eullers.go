package main

import (
	"math/rand"

	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

type rowState struct {
	// cell column: slice index could be an array
	setForCell map[int]int
	// could be a slice maybe?
	cellsInSet map[int][]*cell.BaseCell
	nextSet    int
}

func createEullerState(columns int) *rowState {
	return &rowState{
		setForCell: map[int]int{},
		cellsInSet: map[int][]*cell.BaseCell{},
		nextSet:    0,
	}
}

func (es *rowState) record(set int, c *cell.BaseCell) {
	es.setForCell[c.Column()] = set

	s, ok := es.cellsInSet[set]
	if !ok {

		es.cellsInSet[set] = []*cell.BaseCell{}
	}
	es.cellsInSet[set] = append(s, c)

}

func (es *rowState) setFor(c *cell.BaseCell) int {
	_, ok := es.setForCell[c.Column()]
	if !ok {
		es.record(es.nextSet, c)
		es.nextSet++
	}

	return es.setForCell[c.Column()]
}

func (es *rowState) merge(winner, loser int) {
	cs := es.cellsInSet[loser]
	for i := 0; i < len(cs); i++ {
		c := cs[i]
		es.setForCell[c.Column()] = winner
		es.cellsInSet[winner] = append(es.cellsInSet[winner], c)
	}

	delete(es.cellsInSet, loser)
}

func (es *rowState) next() *rowState {
	return &rowState{
		setForCell: map[int]int{},
		cellsInSet: map[int][]*cell.BaseCell{},
		nextSet:    es.nextSet,
	}
}

func Eullers(sh grid.ShapeHolder) {
	g := sh.GetShape().Grid
	rs := createEullerState(len(g))

	for row := 0; row < len(g); row++ {
		for column := 0; column < len(g[row]); column++ {
			c := g[row][column].(*cell.BaseCell)

			if c.West != nil {
				set := rs.setFor(c)
				priorSet := rs.setFor(c.West)

				// shouldLink := set != priorSet && (c.South == nil || rand.Intn(2) == 0)
				shouldLink := set != priorSet && (c.South == nil || rand.Intn(2) == 0)
				if shouldLink {
					c.Link(c.West)
					rs.merge(priorSet, set)
				}
			}
		}

		rc := g[row][0].(*cell.BaseCell)
		if rc.South != nil {
			nextRow := rs.next()

			for _, cells := range rs.cellsInSet {
				grid.Shuffle(cells)
				for i := 0; i < len(cells); i++ {
					if i == 0 || rand.Intn(3) == 0 {
						cells[i].Link(cells[i].South)
						nextRow.record(rs.setFor(cells[i]), cells[i].South)
					}
				}
			}

			rs = nextRow
		}
	}
}
