package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	pizza := NewVeggieMania()
	pizzaWithCheese := NewCheeseTopping(pizza)
	pizzaWithCheeseAndTomato := NewTomatoTopping(pizzaWithCheese)

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
