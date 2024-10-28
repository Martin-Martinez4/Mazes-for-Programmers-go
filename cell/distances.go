package cell

// PrimeCell will have a weight option that starts at one
// All Cells will have a PrimeCell
// That way weighted distances can be used by all types of cells

type Distances struct {
	root        Cell
	Cells       map[Cell]int
	MaxDistance int
}

const STARTPOINT = 64

func CreateDistances(root Cell) *Distances {

	return &Distances{
		root:  root,
		Cells: map[Cell]int{root: 0},
	}
}

func (d *Distances) Max() (Cell, int) {

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

func WeightedShortestPath(start Cell) *Distances {
	pq := PriorityQueue{}
	pq.Push(start)
	weights := &Distances{root: start, Cells: map[Cell]int{}}

	weights.Cells[start] = 0

	for len(pq) > 0 {
		cell := pq.Pop()

		links := cell.Links()
		for i := 0; i < len(links); i++ {
			neighbor := links[i]
			totalWeight := weights.Cells[cell] + neighbor.Weight()

			val, ok := weights.Cells[neighbor]
			if !ok || totalWeight < val {
				pq.Push(neighbor)
				weights.Cells[neighbor] = totalWeight
			}
		}
	}

	return weights
}

func (d *Distances) longestPath(start Cell) *Distances {
	distances := start.Distances()
	newStart, _ := distances.Max()

	newDistances := newStart.Distances()
	goal, _ := newDistances.Max()

	return newDistances.shortestPath(goal)
}
