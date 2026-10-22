package main

import "fmt"

// Component interface
type BaseCoffee interface {
	Coffee() string
}

// Concrete component
type NormalCoffee struct{}

func (n NormalCoffee) Coffee() string {
	return "Normal coffee"
}

// Decorator: adds sugar
type CoffeeWithSugar struct {
	baseCoffee BaseCoffee
}

func (cs CoffeeWithSugar) Coffee() string {
	return cs.baseCoffee.Coffee() + " + sugar"
}

// Decorator: adds milk
type CoffeeWithMilk struct {
	baseCoffee BaseCoffee
}

func (cm CoffeeWithMilk) Coffee() string {
	return cm.baseCoffee.Coffee() + " + milk"
}

func main() {
	coffee := NormalCoffee{}

	coffeeWithSugar := CoffeeWithSugar{
		baseCoffee: coffee,
	}

	coffeeWithSugarAndMilk := CoffeeWithMilk{
		baseCoffee: coffeeWithSugar,
	}

	fmt.Println(coffee.Coffee())
	fmt.Println(coffeeWithSugar.Coffee())
	fmt.Println(coffeeWithSugarAndMilk.Coffee())
}
