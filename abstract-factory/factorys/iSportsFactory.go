package factorys

import "github.com/dadaxiaoxiao/design-patterns/abstract-factory/products"

type ISportsFactory interface {
	MakeShoe() products.IShoe
	MakeShirt() products.IShirt
}
