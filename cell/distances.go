package cell

type Distances struct {
	root  Cell
	Cells map[Cell]int
}

const STARTPOINT = 64

func CreateDistances(root Cell) *Distances {

	return &Distances{
		root:  root,
		Cells: map[Cell]int{root: 0},
	}
}

func (d *Distances) max() (Cell, int) {

	maxCell := d.root
	maxDistance := 0

	for key, value := range d.Cells {
		distance := value
		cell := key

		if distance > maxDistance {
			maxCell = cell
			maxDistance = distance
		}
	}

	return maxCell, maxDistance
}

func (d *Distances) shortestPath(target Cell) *Distances {
	solutionDistance := CreateDistances(d.root)

	startingCell := d.root
	currentCell := target

	solutionDistance.Cells[currentCell] = d.Cells[currentCell]

	for currentCell != startingCell {

		minDistanceCell := currentCell
		// begin at target
		for _, neighbor := range currentCell.Links() {
			dist := d.Cells[neighbor]

			dist2 := d.Cells[minDistanceCell]
			if dist < dist2 {
				minDistanceCell = neighbor
			}

		}

		if minDistanceCell == currentCell {
			solutionDistance.Cells[startingCell] = STARTPOINT
			return solutionDistance
		} else {
			currentCell = minDistanceCell
			solutionDistance.Cells[currentCell] = d.Cells[currentCell]
		}
	}

	solutionDistance.Cells[startingCell] = STARTPOINT
	return solutionDistance
}

func (d *Distances) longestPath(start Cell) *Distances {
	distances := start.Distances()
	newStart, _ := distances.max()

	newDistances := newStart.Distances()
	goal, _ := newDistances.max()

	return newDistances.shortestPath(goal)
}
