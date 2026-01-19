package main

import (
	"fmt"
)

// Prototype interface
type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

// ConcretePrototypeA
type Car struct {
	model string
	color string
}

func (c *Car) Clone() Prototype {
	clonedCar := *c
	return &clonedCar
}

func (c *Car) GetDetails() string {
	return fmt.Sprintf("Car: Model=%s, Color=%s", c.model, c.color)
}

// ConcretePrototypeB
type Bike struct {
	brand string
	speed string
}

func (b *Bike) Clone() Prototype {
	clonedBike := *b
	return &clonedBike
}

func (b *Bike) GetDetails() string {
	return fmt.Sprintf("Bike: Brand=%s, Speed=%s", b.brand, b.speed)
}

func main() {
	// Create original objects
	originalCar := &Car{model: "BMW", color: "red"}
	originalBike := &Bike{brand: "Harley", speed: "200km/h"}

	// Clone objects
	clonedCar := originalCar.Clone().(*Car)
	clonedBike := originalBike.Clone().(*Bike)

	// Modify cloned objects
	clonedCar.color = "blue"
	clonedBike.speed = "150km/h"

	// Print results
	fmt.Println("Original:", originalCar.GetDetails())
	fmt.Println("Cloned:", clonedCar.GetDetails())
	fmt.Println()
	fmt.Println("Original:", originalBike.GetDetails())
	fmt.Println("Cloned:", clonedBike.GetDetails())
}
