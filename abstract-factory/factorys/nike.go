package factorys

import "github.com/dadaxiaoxiao/design-patterns/abstract-factory/products"

type Nike struct {
}

func (n *Nike) MakeShoe() products.IShoe {
	return &products.NikeShoe{
		Shoe: products.NewShoe("nike", 14),
	}
}

func (n *Nike) MakeShirt() products.IShirt {
	return &products.NikeShirt{
		Shirt: products.NewShirt("nike", 14),
	}
}

func NewNike() ISportsFactory {
	return &Nike{}
}
