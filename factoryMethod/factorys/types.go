package factorys

import "github.com/dadaxiaoxiao/design-patterns/factoryMethod/producuts"

type GunFactory interface {
	CreateGun() producuts.IGun
}
