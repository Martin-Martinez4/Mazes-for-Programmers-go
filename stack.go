package main

type CellStack struct {
	stack []Cell
}

func CreateCellStack() *CellStack {
	return &CellStack{
		stack: []Cell{},
	}
}

func (cs *CellStack) Push(c Cell) {
	cs.stack = append(cs.stack, c)
}

func (cs *CellStack) Pop() Cell {

	length := len(cs.stack)

	if length <= 0 {
		return nil
	}

	cell := cs.stack[length-1]

	cs.stack = cs.stack[:length-1]

	return cell
}

func (cs *CellStack) Peek() Cell {
	return cs.stack[cs.Length()-1]
}

func (cs *CellStack) Length() int {
	return len(cs.stack)
}
