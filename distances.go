package main

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func CreateDistances(root *Cell) *Distances {

	return &Distances{
		root:  root,
		cells: map[*Cell]int{root: 0},
	}
}

func (d *Distances) max() (*Cell, int) {

	maxCell := d.root
	maxDistance := 0

	for key, value := range d.cells {
		distance := value
		cell := key

		if distance > maxDistance {
			maxCell = cell
			maxDistance = distance
		}
	}

	return maxCell, maxDistance
}

func (d *Distances) shortestPath(target *Cell) *Distances {
	solutionDistance := CreateDistances(d.root)

	startingCell := d.root
	currentCell := target

	solutionDistance.cells[currentCell] = d.cells[currentCell]

	for currentCell != startingCell {

		minDistanceCell := currentCell
		// begin at target
		for key := range currentCell.links {
			dist, ok := d.cells[key]
			if ok {

				dist2 := d.cells[minDistanceCell]
				if dist < dist2 {
					minDistanceCell = key
				}

			}

		}

		if minDistanceCell == currentCell {
			solutionDistance.cells[startingCell] = STARTPOINT
			return solutionDistance
		} else {
			currentCell = minDistanceCell
			solutionDistance.cells[currentCell] = d.cells[currentCell]
		}

	}

	solutionDistance.cells[startingCell] = STARTPOINT
	return solutionDistance
}

func (d *Distances) longestPath(start *Cell) *Distances {
	distances := start.distances()
	newStart, _ := distances.max()

	newDistances := newStart.distances()
	goal, _ := newDistances.max()

	return newDistances.shortestPath(goal)
}
