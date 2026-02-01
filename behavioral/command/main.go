package main

import "fmt"

// ============ Command Interface ============

// Command interface - all commands must implement Execute() and Undo()
type Command interface {
	Execute()
	Undo()
}

// ============ Receivers - Objects that perform the actual work ============

// Light receiver
type Light struct {
	location string
	isOn     bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Printf("💡 %s light is ON\n", l.location)
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Printf("🌑 %s light is OFF\n", l.location)
}

// TV receiver
type TV struct {
	isOn      bool
	channel   int
	volume    int
}

func (t *TV) TurnOn() {
	t.isOn = true
	fmt.Println("📺 TV is ON")
}

func (t *TV) TurnOff() {
	t.isOn = false
	fmt.Println("📴 TV is OFF")
}

func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Printf("📺 TV channel set to %d\n", channel)
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
	fmt.Printf("🔊 TV volume set to %d\n", volume)
}

// Fan receiver
type Fan struct {
	speed int // 0=off, 1=low, 2=medium, 3=high
}

func (f *Fan) High() {
	f.speed = 3
	fmt.Println("🌀 Fan is on HIGH speed")
}

func (f *Fan) Medium() {
	f.speed = 2
	fmt.Println("💨 Fan is on MEDIUM speed")
}

func (f *Fan) Low() {
	f.speed = 1
	fmt.Println("🍃 Fan is on LOW speed")
}

func (f *Fan) Off() {
	f.speed = 0
	fmt.Println("⭕ Fan is OFF")
}

// ============ Concrete Commands ============

// LightOnCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// LightOffCommand
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// TVOnCommand
type TVOnCommand struct {
	tv *TV
}

func (c *TVOnCommand) Execute() {
	c.tv.TurnOn()
}

func (c *TVOnCommand) Undo() {
	c.tv.TurnOff()
}

// FanHighCommand
type FanHighCommand struct {
	fan       *Fan
	prevSpeed int
}

func (c *FanHighCommand) Execute() {
	c.prevSpeed = c.fan.speed
	c.fan.High()
}

func (c *FanHighCommand) Undo() {
	switch c.prevSpeed {
	case 3:
		c.fan.High()
	case 2:
		c.fan.Medium()
	case 1:
		c.fan.Low()
	default:
		c.fan.Off()
	}
}

// MacroCommand - executes multiple commands
type MacroCommand struct {
	commands []Command
}

func (c *MacroCommand) Execute() {
	for _, cmd := range c.commands {
		cmd.Execute()
	}
}

func (c *MacroCommand) Undo() {
	// Undo in reverse order
	for i := len(c.commands) - 1; i >= 0; i-- {
		c.commands[i].Undo()
	}
}

// ============ Invoker - Remote Control ============

type RemoteControl struct {
	commands    [7]Command // 7 buttons
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	rc := &RemoteControl{}
	// Initialize with NoCommand (Null Object pattern)
	noCmd := &NoCommand{}
	for i := 0; i < 7; i++ {
		rc.commands[i] = noCmd
	}
	rc.undoCommand = noCmd
	return rc
}

func (rc *RemoteControl) SetCommand(slot int, command Command) {
	rc.commands[slot] = command
}

func (rc *RemoteControl) PressButton(slot int) {
	rc.commands[slot].Execute()
	rc.undoCommand = rc.commands[slot]
}

func (rc *RemoteControl) PressUndo() {
	rc.undoCommand.Undo()
}

// NoCommand - Null Object pattern
type NoCommand struct{}

func (c *NoCommand) Execute() {}
func (c *NoCommand) Undo()    {}

func main() {
	fmt.Println("===== Command Pattern Example: Smart Home Remote Control =====\n")

	// Create receivers (devices)
	livingRoomLight := &Light{location: "Living Room"}
	kitchenLight := &Light{location: "Kitchen"}
	tv := &TV{}
	fan := &Fan{}

	// Create commands
	livingRoomLightOn := &LightOnCommand{light: livingRoomLight}
	livingRoomLightOff := &LightOffCommand{light: livingRoomLight}
	kitchenLightOn := &LightOnCommand{light: kitchenLight}
	tvOn := &TVOnCommand{tv: tv}
	fanHigh := &FanHighCommand{fan: fan}

	// Create remote control (invoker)
	remote := NewRemoteControl()

	// Set commands to buttons
	remote.SetCommand(0, livingRoomLightOn)
	remote.SetCommand(1, livingRoomLightOff)
	remote.SetCommand(2, kitchenLightOn)
	remote.SetCommand(3, tvOn)
	remote.SetCommand(4, fanHigh)

	// Test individual commands
	fmt.Println("--- Testing Individual Commands ---")
	fmt.Println("Press button 0:")
	remote.PressButton(0) // Living room light on

	fmt.Println("\nPress button 4:")
	remote.PressButton(4) // Fan high

	fmt.Println("\nPress Undo:")
	remote.PressUndo() // Undo fan (turn off)

	fmt.Println("\nPress button 3:")
	remote.PressButton(3) // TV on

	fmt.Println("\nPress Undo:")
	remote.PressUndo() // Undo TV (turn off)

	// Test Macro Command (Party Mode!)
	fmt.Println("\n--- Testing Macro Command (Party Mode) ---")
	partyMode := &MacroCommand{
		commands: []Command{
			livingRoomLightOn,
			kitchenLightOn,
			tvOn,
			fanHigh,
		},
	}

	fmt.Println("Activating Party Mode:")
	partyMode.Execute()

	fmt.Println("\nDeactivating Party Mode (Undo):")
	partyMode.Undo()

	fmt.Println("\n✅ Command Pattern Demo Complete!")
}
