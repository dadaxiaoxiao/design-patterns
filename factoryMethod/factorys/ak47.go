package factorys

import "github.com/dadaxiaoxiao/design-patterns/factoryMethod/producuts"

type Ak47Factory struct {
}

func (a *Ak47Factory) CreateGun() producuts.IGun {
	return producuts.NewAk47()
}

func NewAk47Factory() GunFactory {
	return &Ak47Factory{}
}
