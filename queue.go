package main

type CellQueue struct {
	elements chan interface{}
}

func NewQueue(size int) *CellQueue {
	return &CellQueue{
		elements: make(chan interface{}, size),
	}
}

func (queue *CellQueue) Push(element interface{}) {
	select {
	case queue.elements <- element:
	default:
		panic("Queue full")
	}
}

func (queue *CellQueue) Pop() interface{} {
	select {
	case e := <-queue.elements:
		return e
	default:
		panic("Queue empty")
	}
}
