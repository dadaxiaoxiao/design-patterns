package abstract_factory

import (
	"fmt"
	"github.com/dadaxiaoxiao/design-patterns/abstract-factory/factorys"
	"github.com/dadaxiaoxiao/design-patterns/abstract-factory/products"
	"testing"
)

func Test(t *testing.T) {
	adidasFactory := factorys.NewAdidas()
	nikeFactory := factorys.NewNike()

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}

func printShoeDetails(s products.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s products.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
