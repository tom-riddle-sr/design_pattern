package main

import "fmt"

type Request struct {
	Name    string
	Age     int
	Country string
	Amount  float64
}

type Handler interface {
	SetNext(next Handler) Handler
	Handle(req Request)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) Handler {
	b.next = next
	return next
}

func (b *BaseHandler) Handle(req Request) {
	if b.next != nil {
		b.next.Handle(req)
		return
	}
	fmt.Printf("Request rejected for %s\n", req.Name)
}

type AgeCheck struct{ BaseHandler }

func (h *AgeCheck) Handle(req Request) {
	if req.Age < 18 {
		fmt.Printf("AgeCheck failed for %s (age %d)\n", req.Name, req.Age)
		return
	}
	h.BaseHandler.Handle(req)
}

type CountryCheck struct{ BaseHandler }

func (h *CountryCheck) Handle(req Request) {
	allowed := map[string]bool{"US": true, "TW": true, "JP": true}
	if !allowed[req.Country] {
		fmt.Printf("CountryCheck failed for %s (country %s)\n", req.Name, req.Country)
		return
	}
	h.BaseHandler.Handle(req)
}

type BalanceCheck struct{ BaseHandler }

func (h *BalanceCheck) Handle(req Request) {
	if req.Amount > 1000 {
		fmt.Printf("BalanceCheck failed for %s (amount %.2f)\n", req.Name, req.Amount)
		return
	}
	fmt.Printf("Request approved for %s (amount %.2f)\n", req.Name, req.Amount)
}

func main() {
	fmt.Println("===== Chain of Responsibility (Verification) =====")

	age := &AgeCheck{}
	country := &CountryCheck{}
	balance := &BalanceCheck{}

	age.SetNext(country).SetNext(balance)

	requests := []Request{
		{Name: "Amy", Age: 17, Country: "TW", Amount: 200},
		{Name: "Ben", Age: 25, Country: "CN", Amount: 200},
		{Name: "Cory", Age: 30, Country: "US", Amount: 1500},
		{Name: "Dana", Age: 22, Country: "JP", Amount: 300},
	}

	for _, r := range requests {
		age.Handle(r)
	}
}
