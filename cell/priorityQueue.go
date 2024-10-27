package cell

import (
	"fmt"
	"sort"
)

// Implemented using a heap as shown here: https://pkg.go.dev/container/heap#example__priorityQueue

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// sort lowest weight last because there will be a lot of Poping
	return pq[i].Weight() > pq[j].Weight()
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x Cell) {
	weight := x.Weight()

	if len(*pq) == 0 {
		*pq = append(*pq, x)
		return
	}

	// check both ends first
	if weight <= (*pq)[len(*pq)-1].Weight() {
		*pq = append(*pq, x)
	} else if weight >= (*pq)[0].Weight() {
		*pq = append([]Cell{x}, *pq...)
	} else {

		index := len(*pq) - 1
		for (*pq)[index].Weight() < weight {
			index--

			if index < 0 {
				return
			}
		}

		fmt.Println("index: ", index)
		*pq = append(*pq, (*pq)[len(*pq)-1])
		copy((*pq)[index+1:], (*pq)[index:len(*pq)-1])
		(*pq)[index+1] = x

	}
}

func (pq *PriorityQueue) Pop() Cell {
	if len(*pq) == 0 {
		return nil
	}
	p := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return p
}

func (pq *PriorityQueue) Peek() Cell {

	return (*pq)[len(*pq)-1]
}

// Takes a long time but should never really be used
func (pq *PriorityQueue) Repair(index int) {
	// Just a wrap for sorting the slice
	sort.Sort(pq)

}
