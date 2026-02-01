package main

import "fmt"

// Memento stores the state
type Memento struct {
	state string
}

// Originator creates and restores mementos
type Originator struct {
	state string
}

func (o *Originator) SetState(s string) {
	o.state = s
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) Save() Memento {
	return Memento{state: o.state}
}

func (o *Originator) Restore(m Memento) {
	o.state = m.state
}

// Caretaker manages mementos
type Caretaker struct {
	history []Memento
}

func (c *Caretaker) Backup(m Memento) {
	c.history = append(c.history, m)
}

func (c *Caretaker) Undo() (Memento, bool) {
	if len(c.history) == 0 {
		return Memento{}, false
	}
	m := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]
	return m, true
}

func main() {
	fmt.Println("===== Memento Pattern Example =====")

	origin := &Originator{}
	caretaker := &Caretaker{}

	origin.SetState("State #1")
	caretaker.Backup(origin.Save())
	fmt.Println("Current state:", origin.GetState())

	origin.SetState("State #2")
	caretaker.Backup(origin.Save())
	fmt.Println("Current state:", origin.GetState())

	origin.SetState("State #3")
	fmt.Println("Current state:", origin.GetState())

	// Undo
	if m, ok := caretaker.Undo(); ok {
		origin.Restore(m)
		fmt.Println("Undo! State restored to:", origin.GetState())
	}
	if m, ok := caretaker.Undo(); ok {
		origin.Restore(m)
		fmt.Println("Undo! State restored to:", origin.GetState())
	}
}
