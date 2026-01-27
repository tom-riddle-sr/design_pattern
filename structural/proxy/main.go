package main

import "fmt"

// Subject interface
type Subject interface {
	Request()
}

// RealSubject: The real object that the proxy represents
type RealSubject struct{}

func (r *RealSubject) Request() {
	fmt.Println("RealSubject: Handling request.")
}

// Proxy: The proxy object that controls access to the RealSubject
type Proxy struct {
	realSubject *RealSubject
}

func NewProxy(realSubject *RealSubject) *Proxy {
	return &Proxy{realSubject: realSubject}
}

func (p *Proxy) Request() {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy: Logging before forwarding the request.")
	p.realSubject.Request()
	fmt.Println("Proxy: Logging after forwarding the request.")
}

func main() {
	fmt.Println("===== Proxy Pattern Example =====")

	// Create a RealSubject
	realSubject := &RealSubject{}

	// Create a Proxy for the RealSubject
	proxy := NewProxy(realSubject)

	// Use the Proxy to make a request
	proxy.Request()
}
