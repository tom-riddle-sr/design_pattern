package main

import "fmt"

// ===== Component Interface =====
type Coffee interface {
	GetDescription() string
	GetCost() float64
}

// ===== Concrete Component =====
// 基礎咖啡
type SimpleCoffee struct{}

func (sc *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

func (sc *SimpleCoffee) GetCost() float64 {
	return 100.0
}

// ===== Decorators =====
// 裝飾者基礎結構
type CoffeeDecorator struct {
	coffee Coffee
}

// 添加牛奶的裝飾者
type MilkDecorator struct {
	CoffeeDecorator
}

func (md *MilkDecorator) GetDescription() string {
	return md.coffee.GetDescription() + " + Milk"
}

func (md *MilkDecorator) GetCost() float64 {
	return md.coffee.GetCost() + 50.0
}

func NewMilkDecorator(coffee Coffee) Coffee {
	return &MilkDecorator{CoffeeDecorator{coffee: coffee}}
}

// 添加糖的裝飾者
type SugarDecorator struct {
	CoffeeDecorator
}

func (sd *SugarDecorator) GetDescription() string {
	return sd.coffee.GetDescription() + " + Sugar"
}

func (sd *SugarDecorator) GetCost() float64 {
	return sd.coffee.GetCost() + 10.0
}

func NewSugarDecorator(coffee Coffee) Coffee {
	return &SugarDecorator{CoffeeDecorator{coffee: coffee}}
}

// 添加鮮奶油的裝飾者
type WhipDecorator struct {
	CoffeeDecorator
}

func (wd *WhipDecorator) GetDescription() string {
	return wd.coffee.GetDescription() + " + Whip Cream"
}

func (wd *WhipDecorator) GetCost() float64 {
	return wd.coffee.GetCost() + 20.0
}

func NewWhipDecorator(coffee Coffee) Coffee {
	return &WhipDecorator{CoffeeDecorator{coffee: coffee}}
}

// 添加巧克力的裝飾者
type ChocolateDecorator struct {
	CoffeeDecorator
}

func (cd *ChocolateDecorator) GetDescription() string {
	return cd.coffee.GetDescription() + " + Chocolate"
}

func (cd *ChocolateDecorator) GetCost() float64 {
	return cd.coffee.GetCost() + 30.0
}

func NewChocolateDecorator(coffee Coffee) Coffee {
	return &ChocolateDecorator{CoffeeDecorator{coffee: coffee}}
}

func main() {
	fmt.Println("===== Coffee Decorator Pattern Example =====\n")

	// 創建基礎咖啡
	simpleCoffee := &SimpleCoffee{}
	displayCoffee(simpleCoffee)

	// 添加牛奶
	coffeeWithMilk := NewMilkDecorator(simpleCoffee)
	displayCoffee(coffeeWithMilk)

	// 添加牛奶和糖
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
	displayCoffee(coffeeWithMilkAndSugar)

	// 添加牛奶、糖和鮮奶油
	deluxeCoffee := NewWhipDecorator(coffeeWithMilkAndSugar)
	displayCoffee(deluxeCoffee)

	// 另一種組合：基礎咖啡 + 巧克力 + 鮮奶油
	fmt.Println("\n===== Different Combination =====\n")
	chocolateCoffee := NewChocolateDecorator(simpleCoffee)
	coffeeWithChocolateAndWhip := NewWhipDecorator(chocolateCoffee)
	displayCoffee(coffeeWithChocolateAndWhip)
}

func displayCoffee(coffee Coffee) {
	fmt.Printf("☕ %s\n", coffee.GetDescription())
	fmt.Printf("   💰 Cost: ¥%.2f\n\n", coffee.GetCost())
}
