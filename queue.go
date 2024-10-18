package main

type CellQueue struct {
	elements chan *Cell
}

func CreateQueue(size int) *CellQueue {
	return &CellQueue{
		elements: make(chan *Cell, size),
	}
}

func (queue *CellQueue) Push(element *Cell) {
	select {
	case queue.elements <- element:
	default:
		panic("Queue full")
	}
}

func (queue *CellQueue) Pop() *Cell {
	select {
	case e := <-queue.elements:
		return e
	default:
		return nil
	}
}
