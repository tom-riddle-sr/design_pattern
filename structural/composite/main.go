package main

import "fmt"

type Component interface {
	ShoppingItem()
}

type Cart struct {
	Bundles []Bundles
}

func (c *Cart) ShoppingItem() {
	fmt.Println("🛒 Shopping Cart Contents:")
	for _, bundle := range c.Bundles {
		bundle.ShoppingItem()
	}
}

func (c *Cart) GetTotalPrice() float64 {
	total := 0.0
	for _, bundle := range c.Bundles {
		total += bundle.GetTotalPrice()
	}
	return total
}

type Bundles struct {
	Name    string
	Price   float64
	Product []Products
}

func (b *Bundles) ShoppingItem() {
	fmt.Printf("📦 Bundle: %s (¥%.2f)\n", b.Name, b.Price)
	for _, product := range b.Product {
		fmt.Printf("  📦 %s\n", product.Name)
	}
}

func (b *Bundles) GetTotalPrice() float64 {
	total := b.Price
	for _, product := range b.Product {
		total += product.Price
	}
	return total
}

type Products struct {
	Name  string
	Price float64
}

func (p *Products) ShoppingItem() {
	fmt.Printf("Product Name: %s, Price: %.2f\n", p.Name, p.Price)
}

func main() {
	cart := &Cart{}

	// 創建電子產品套餐
	electronicBundle := Bundles{
		Name:  "Electronic Bundle",
		Price: 0,
		Product: []Products{
			{Name: "Phone", Price: 1500.0},
			{Name: "Charger", Price: 500.0},
		},
	}

	// 創建日用品套餐
	dailyBundle := Bundles{
		Name:  "Daily Bundle",
		Price: 0,
		Product: []Products{
			{Name: "Toothpaste", Price: 50.0},
			{Name: "Soap", Price: 100.0},
		},
	}

	// 添加套餐到購物車
	cart.Bundles = append(cart.Bundles, electronicBundle)
	cart.Bundles = append(cart.Bundles, dailyBundle)
	cart.Bundles = append(cart.Bundles, Bundles{
		Name:    "Books",
		Price:   100.0,
		Product: []Products{},
	})

	// 顯示購物車
	cart.ShoppingItem()

	// 顯示總價
	fmt.Printf("\n💰 Total Price: ¥%.2f\n", cart.GetTotalPrice())
}
