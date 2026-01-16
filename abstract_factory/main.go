package main

type Drink interface {
	Serve() string
	AddIce() string
	SetTemp() string
}

type Coffee struct{}

func (c *Coffee) Serve() string {
	return "Serving a cup of Coffee"
}

func (c *Coffee) AddIce() string {
	return "Adding ice to Coffee"
}

func (c *Coffee) SetTemp() string {
	return "Setting Coffee temperature to hot"
}

type Tea struct{}

func (t *Tea) Serve() string {
	return "Serving a cup of Tea"
}

func (t *Tea) AddIce() string {
	return "Adding ice to Tea"
}

func (t *Tea) SetTemp() string {
	return "Setting Tea temperature to warm"
}

type DrinkFactory interface {
	CreateDrink() Drink
}

type CoffeeFactory struct{}

func (f *CoffeeFactory) CreateDrink() Drink {
	return &Coffee{}
}

type TeaFactory struct{}

func (f *TeaFactory) CreateDrink() Drink {
	return &Tea{}
}

func GetDrinkFactory(drinkType string) DrinkFactory {
	if drinkType == "coffee" {
		return &CoffeeFactory{}
	} else if drinkType == "tea" {
		return &TeaFactory{}
	}
	return nil
}

func main() {
	coffeeFactory := GetDrinkFactory("coffee")
	coffee := coffeeFactory.CreateDrink()
	println(coffee.Serve())
	println(coffee.AddIce())
	println(coffee.SetTemp())

	teaFactory := GetDrinkFactory("tea")
	tea := teaFactory.CreateDrink()
	println(tea.Serve())
	println(tea.AddIce())
	println(tea.SetTemp())
}
