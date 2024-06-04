package decorator

type CheeseTopping struct {
	pizza IPizza
}

func NewCheeseTopping(pizza IPizza) IPizza {
	return &TomatoTopping{
		pizza: pizza,
	}
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
