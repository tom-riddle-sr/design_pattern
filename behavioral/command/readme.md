# **Command Pattern**

**Command** is a behavioral design pattern that encapsulates a request as an object, allowing you to parameterize clients with different requests, queue requests, log them, and support undoable operations.

## Key Concepts

- **Command Interface**: Declares the execution interface (usually `Execute()`)
- **Concrete Command**: Implements the command and holds a reference to the receiver
- **Receiver**: The object that performs the actual work
- **Invoker**: Asks the command to execute the request
- **Client**: Creates and configures concrete command objects

## Diagram

```
Client
   ↓
   creates Command
   ↓
Invoker → Execute() → ConcreteCommand → Receiver.Action()
                           ↓
                      stores receiver

Command Interface
   ↑
   ├── TurnOnCommand
   ├── TurnOffCommand
   └── DimLightCommand

Example: Remote control with buttons
         Button → Execute() → TurnOnCommand → Light.TurnOn()
```

## Advantages

- **Decoupling**: Separates the object that invokes the operation from the object that performs it
- **Undo/Redo**: Easy to implement by storing command history
- **Queueing**: Commands can be queued and executed at different times
- **Logging**: Command history can be logged for auditing or replay
- **Macro Commands**: Multiple commands can be composed into one
- **Open/Closed Principle**: New commands can be added without changing existing code

## Disadvantages

- **Increased Classes**: Each command requires a separate class
- **Complexity**: Can add unnecessary complexity for simple operations
- **Memory Usage**: Storing command history consumes memory

## Real-World Examples

- **GUI Buttons/Menu Items** - Each button executes a command (Copy, Paste, Save)
- **Remote Control** - Each button is a command (TurnOn, TurnOff, ChangeChannel)
- **Text Editor Undo/Redo** - Each edit operation is a command
- **Task Scheduler** - Queue commands to execute later
- **Transaction Systems** - Each transaction is a command that can be rolled back
- **Game Input** - Map keyboard/controller inputs to commands

## Comparison with Other Patterns

| Pattern        | Purpose                          | Example                    |
|--------------- |----------------------------------|----------------------------|
| **Command**    | Encapsulate requests as objects  | Button actions, Undo/Redo  |
| **Strategy**   | Encapsulate algorithms           | Sorting, Payment methods   |
| **State**      | Change behavior based on state   | Order status transitions   |
| **Memento**    | Save/restore state               | Checkpoints, Snapshots     |

## Characteristics

- **Request as Object**: Encapsulates a request as an independent object
- **Decoupling**: Invoker doesn't know about the receiver or how the action is performed
- **Parameterization**: Objects can be parameterized with different commands
- **Reversibility**: Commands can implement Undo() to reverse their effects
- **Command History**: Can store executed commands for replay, audit, or undo

## When to Use

✅ Need to parameterize objects with operations
✅ Need to queue, schedule, or execute requests at different times
✅ Need to support Undo/Redo operations
✅ Need to log changes for auditing or replay
✅ Need to support macro commands (compose multiple commands)
❌ Avoid for simple, one-to-one method calls (adds unnecessary complexity)

## Example Scenarios

- Text editor with Undo/Redo functionality
- Remote control for home automation devices
- Job queue/task scheduler systems
- Transaction management in databases
- Game input mapping (keyboard to actions)