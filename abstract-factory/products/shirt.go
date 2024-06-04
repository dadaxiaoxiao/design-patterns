package products

type Shirt struct {
	logo string
	size int
}

func NewShirt(logo string, size int) *Shirt {
	return &Shirt{
		logo: logo,
		size: size,
	}
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}
