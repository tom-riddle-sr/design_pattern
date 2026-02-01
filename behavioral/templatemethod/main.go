package main

import "fmt"

// Beverage is the abstract template that defines the algorithm structure
type Beverage interface {
	// Template Method - defines the algorithm skeleton
	PrepareRecipe()

	// Primitive operations - must be implemented by concrete classes
	Brew()
	AddCondiments()

	// Common operations - shared by all beverages
	boilWater()
	pourInCup()
}

// BaseBeverage provides common implementation
type BaseBeverage struct {
	beverage Beverage
}

func (b *BaseBeverage) PrepareRecipe() {
	b.beverage.boilWater()
	b.beverage.Brew()
	b.beverage.pourInCup()
	b.beverage.AddCondiments()
}

func (b *BaseBeverage) boilWater() {
	fmt.Println("💧 Boiling water...")
}

func (b *BaseBeverage) pourInCup() {
	fmt.Println("☕ Pouring into cup...")
}

// ============ Concrete Implementations ============

// Tea concrete implementation
type Tea struct {
	BaseBeverage
}

func NewTea() *Tea {
	t := &Tea{}
	t.beverage = t
	return t
}

func (t *Tea) Brew() {
	fmt.Println("🍵 Steeping the tea leaves...")
}

func (t *Tea) AddCondiments() {
	fmt.Println("🍋 Adding lemon...")
}

// Coffee concrete implementation
type Coffee struct {
	BaseBeverage
}

func NewCoffee() *Coffee {
	c := &Coffee{}
	c.beverage = c
	return c
}

func (c *Coffee) Brew() {
	fmt.Println("☕ Dripping coffee through filter...")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("🥛 Adding sugar and milk...")
}

// HotChocolate concrete implementation
type HotChocolate struct {
	BaseBeverage
}

func NewHotChocolate() *HotChocolate {
	h := &HotChocolate{}
	h.beverage = h
	return h
}

func (h *HotChocolate) Brew() {
	fmt.Println("🍫 Melting chocolate powder...")
}

func (h *HotChocolate) AddCondiments() {
	fmt.Println("🍰 Adding whipped cream and marshmallows...")
}

func main() {
	fmt.Println("===== Template Method Pattern Example: Beverage Preparation =====\n")

	// Prepare Tea
	fmt.Println("--- Making Tea ---")
	tea := NewTea()
	tea.PrepareRecipe()
	fmt.Println()

	// Prepare Coffee
	fmt.Println("--- Making Coffee ---")
	coffee := NewCoffee()
	coffee.PrepareRecipe()
	fmt.Println()

	// Prepare Hot Chocolate
	fmt.Println("--- Making Hot Chocolate ---")
	hotChocolate := NewHotChocolate()
	hotChocolate.PrepareRecipe()
	fmt.Println()

	fmt.Println("✅ All beverages prepared using the same template method!")
}
