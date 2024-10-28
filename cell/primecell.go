package cell

// Prime short for primordial aka the first
type PrimeCell struct {
	row    int
	column int
	weight int
}

func CreatePrimeCell(row, column int) *PrimeCell {
	return &PrimeCell{row: row, column: column, weight: 1}
}

func (p *PrimeCell) Row() int {
	return p.row
}

func (p *PrimeCell) Column() int {
	return p.column
}

func (p *PrimeCell) Weight() int {
	return p.weight
}

func (p *PrimeCell) SetWeight(weight int) {
	p.weight = weight
}
