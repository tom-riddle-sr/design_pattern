package main

import "fmt"

// ============ Strategy Interface ============

// PaymentStrategy defines the interface for payment methods
type PaymentStrategy interface {
	Pay(amount float64)
}

// ============ Concrete Strategies ============

// CreditCardPayment is a concrete strategy for credit card payments
type CreditCardPayment struct {
	cardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{cardNumber: cardNumber}
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card (Card Number: %s)\n", amount, c.cardNumber)
}

// PayPalPayment is a concrete strategy for PayPal payments
type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email: email}
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal (Email: %s)\n", amount, p.email)
}

// ============ Context ============

// PaymentContext is the context that uses a PaymentStrategy
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	p.strategy.Pay(amount)
}

func main() {
	fmt.Println("===== Strategy Pattern Example: Payment System =====\n")

	// Create payment strategies
	creditCard := NewCreditCardPayment("1234-5678-9012-3456")
	payPal := NewPayPalPayment("user@example.com")

	// Create payment context
	payment := NewPaymentContext(creditCard)

	// Pay using Credit Card
	fmt.Println("--- Paying with Credit Card ---")
	payment.Pay(100.0)

	// Change strategy to PayPal
	fmt.Println("\n--- Switching to PayPal ---")
	payment.SetStrategy(payPal)

	// Pay using PayPal
	payment.Pay(200.0)

	fmt.Println("\n--- Benefits of Strategy Pattern ---")
	fmt.Println("✅ Allows dynamic selection of algorithms at runtime")
	fmt.Println("✅ Promotes code reusability and flexibility")
	fmt.Println("✅ Adheres to the Open/Closed Principle")
}
