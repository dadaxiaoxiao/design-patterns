package factorys

import "github.com/dadaxiaoxiao/design-patterns/factoryMethod/producuts"

type DemoFactory struct {
}

func (d *DemoFactory) CreateGun() producuts.IGun {
	return producuts.NewDemo()
}

func NewDemoFactory() GunFactory {
	return &DemoFactory{}
}
