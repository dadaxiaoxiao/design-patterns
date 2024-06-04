package decorator

type VeggieMania struct {
}

func NewVeggieMania() IPizza {
	return &VeggieMania{}
}

func (p *VeggieMania) getPrice() int {
	return 15
}
