package decorator

type TomatoTopping struct {
	pizza IPizza
}

func NewTomatoTopping(pizza IPizza) IPizza {
	return &TomatoTopping{
		pizza: pizza,
	}
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
