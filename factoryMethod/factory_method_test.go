package factoryMethod

import (
	"fmt"
	"github.com/dadaxiaoxiao/design-patterns/factoryMethod/factorys"
	"github.com/dadaxiaoxiao/design-patterns/factoryMethod/producuts"
	"testing"
)

// getFactory 这里作为客户不得不耦合 type
func getFactory(factoryType string) (factorys.GunFactory, error) {
	if factoryType == "ak47" {
		return factorys.NewAk47Factory(), nil
	}
	if factoryType == "demo" {
		return factorys.NewDemoFactory(), nil
	}
	return nil, fmt.Errorf("Wrong factory type passed")
}

func TestFactoryMethod(t *testing.T) {

	ak47, _ := getFactory("ak47")
	demo, _ := getFactory("demo")

	printDetails(ak47.CreateGun())
	printDetails(demo.CreateGun())

}

func printDetails(g producuts.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
