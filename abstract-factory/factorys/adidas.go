package factorys

import "github.com/dadaxiaoxiao/design-patterns/abstract-factory/products"

type Adidas struct {
}

func NewAdidas() ISportsFactory {
	return &Adidas{}
}

func (a *Adidas) MakeShoe() products.IShoe {
	return &products.AdidasShoe{
		Shoe: products.NewShoe("adidas", 14),
	}
}

func (a *Adidas) MakeShirt() products.IShirt {
	return &products.AdidasShirt{
		Shirt: products.NewShirt("adidas", 14),
	}
}
