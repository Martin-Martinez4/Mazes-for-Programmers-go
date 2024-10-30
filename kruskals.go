package main

import (
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/cell"
	"github.com/Martin-Martinez4/Mazes-for-Programmers-go/grid"
)

// Skipped weaved Kruskals because it would be too much of a hassel with how the code is structured.
type Set map[*Set][]*cell.BaseCell

type state struct {
	neighbors  [][]*cell.BaseCell
	setForCell map[*cell.BaseCell]*Set
	cellsInSet *Set
}

func (s *state) initialize(sh grid.ShapeHolder) {
	grid := sh.GetShape().Grid
	s.setForCell = map[*cell.BaseCell]*Set{}
	s.cellsInSet = &Set{}

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			set := &Set{}
			c := grid[row][column].(*cell.BaseCell)

			s.setForCell[c] = set
			(*s.cellsInSet)[set] = []*cell.BaseCell{c}

			if c.South != nil {
				s.neighbors = append(s.neighbors, []*cell.BaseCell{c, c.South})
			}
			if c.East != nil {
				s.neighbors = append(s.neighbors, []*cell.BaseCell{c, c.East})
			}
		}
	}
}

func (s *state) CanMerge(left, right *cell.BaseCell) bool {
	return s.setForCell[left] != s.setForCell[right]
}

func (s *state) Merge(left *cell.BaseCell, right *cell.BaseCell) {
	left.Link(right)
	winner := s.setForCell[left]
	loser := s.setForCell[right]

	losers := (*s.cellsInSet)[loser]

	for _, c := range losers {
		(*s.cellsInSet)[winner] = append((*s.cellsInSet)[winner], c)
		s.setForCell[c] = winner
	}
	delete(*s.cellsInSet, loser)
}

func Kruskals(sh grid.ShapeHolder) {
	s := &state{}
	s.initialize(sh)

	grid.Shuffle(s.neighbors)

	for len(s.neighbors) > 0 {
		pair := s.neighbors[len(s.neighbors)-1]
		s.neighbors = s.neighbors[:len(s.neighbors)-1]
		left, right := pair[0], pair[1]
		if s.CanMerge(left, right) {
			s.Merge(left, right)
		}
	}
}
