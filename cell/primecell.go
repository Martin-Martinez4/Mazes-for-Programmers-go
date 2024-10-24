package cell

// Prime short for primordial aka the first
type PrimeCell struct {
	row    int
	column int
}

func (p *PrimeCell) Row() int {
	return p.row
}

func (p *PrimeCell) Column() int {
	return p.column
}
