package main

import (
	"fmt"
	"time"
)

// State interface - defines what actions can be performed
type State interface {
	InsertMoney(v *VendingMachine)
	PressButton(v *VendingMachine)
	Dispense(v *VendingMachine)
}

// VendingMachine (Context) - holds current state
type VendingMachine struct {
	currentState State
	itemCount    int
}

func NewVendingMachine(itemCount int) *VendingMachine {
	v := &VendingMachine{itemCount: itemCount}
	v.SetState(&NoMoneyState{})
	return v
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) InsertMoney() {
	v.currentState.InsertMoney(v)
}

func (v *VendingMachine) PressButton() {
	v.currentState.PressButton(v)
}

func (v *VendingMachine) Dispense() {
	v.currentState.Dispense(v)
}

func (v *VendingMachine) GetItemCount() int {
	return v.itemCount
}

func (v *VendingMachine) DecreaseItem() {
	if v.itemCount > 0 {
		v.itemCount--
	}
}

// ============ Concrete States ============

// NoMoneyState - waiting for money
type NoMoneyState struct{}

func (s *NoMoneyState) InsertMoney(v *VendingMachine) {
	fmt.Println("💰 Money inserted! Please select an item.")
	v.SetState(&HasMoneyState{})
}

func (s *NoMoneyState) PressButton(v *VendingMachine) {
	fmt.Println("❌ Please insert money first!")
}

func (s *NoMoneyState) Dispense(v *VendingMachine) {
	fmt.Println("❌ Please insert money first!")
}

// HasMoneyState - money received, waiting for selection
type HasMoneyState struct{}

func (s *HasMoneyState) InsertMoney(v *VendingMachine) {
	fmt.Println("⚠️  You already inserted money!")
}

func (s *HasMoneyState) PressButton(v *VendingMachine) {
	if v.GetItemCount() > 0 {
		fmt.Println("✅ Item selected! Dispensing...")
		v.SetState(&DispensingState{})
		v.Dispense()
	} else {
		fmt.Println("❌ Sorry, out of stock! Returning money...")
		v.SetState(&NoMoneyState{})
	}
}

func (s *HasMoneyState) Dispense(v *VendingMachine) {
	fmt.Println("⚠️  Please press the button first!")
}

// DispensingState - item is being dispensed
type DispensingState struct{}

func (s *DispensingState) InsertMoney(v *VendingMachine) {
	fmt.Println("⏳ Please wait, dispensing in progress...")
}

func (s *DispensingState) PressButton(v *VendingMachine) {
	fmt.Println("⏳ Please wait, dispensing in progress...")
}

func (s *DispensingState) Dispense(v *VendingMachine) {
	fmt.Println("🔄 Dispensing item...")
	time.Sleep(1 * time.Second)
	v.DecreaseItem()
	fmt.Printf("🎉 Item dispensed! Remaining items: %d\n", v.GetItemCount())

	if v.GetItemCount() > 0 {
		v.SetState(&NoMoneyState{})
	} else {
		fmt.Println("📦 Machine is now empty!")
		v.SetState(&OutOfStockState{})
	}
}

// OutOfStockState - no items available
type OutOfStockState struct{}

func (s *OutOfStockState) InsertMoney(v *VendingMachine) {
	fmt.Println("❌ Out of stock! Cannot accept money.")
}

func (s *OutOfStockState) PressButton(v *VendingMachine) {
	fmt.Println("❌ Out of stock!")
}

func (s *OutOfStockState) Dispense(v *VendingMachine) {
	fmt.Println("❌ Out of stock!")
}

func main() {
	fmt.Println("===== State Pattern Example: Vending Machine =====\n")

	// Create vending machine with 2 items
	machine := NewVendingMachine(2)

	// Test scenario 1: Normal purchase
	fmt.Println("--- Scenario 1: Normal Purchase ---")
	machine.PressButton()      // No money
	machine.InsertMoney()      // Insert money
	machine.InsertMoney()      // Try to insert again
	machine.PressButton()      // Select item
	fmt.Println()

	// Test scenario 2: Another purchase
	fmt.Println("--- Scenario 2: Second Purchase ---")
	machine.InsertMoney()
	machine.PressButton()
	fmt.Println()

	// Test scenario 3: Out of stock
	fmt.Println("--- Scenario 3: Out of Stock ---")
	machine.InsertMoney()      // Try to insert money
	machine.PressButton()      // Try to select item
}
