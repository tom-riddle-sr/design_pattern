# **State Pattern**

**State** is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.

## Key Concepts

- **Context**: The object whose behavior changes based on its state
- **State Interface**: Defines the interface for all concrete states
- **Concrete States**: Each state implements specific behavior for that state

## Diagram

```
Context
   ↓
   state: State
   ↓
   request() → state.handle()

State Interface
   ↑
   ├── ConcreteStateA
   ├── ConcreteStateB
   └── ConcreteStateC

Example: Document workflow (Draft → Review → Published)
```

## Advantages

- **Single Responsibility**: Each state class handles one state's behavior
- **Open/Closed Principle**: Easy to add new states without modifying existing code
- **Eliminates Complex Conditionals**: Replaces large if-else/switch statements
- **Clean State Transitions**: State transitions are explicit and manageable

## Disadvantages

- **Increased Classes**: Each state requires a separate class
- **Overkill for Simple Cases**: May be too complex for objects with few states

## Real-World Examples

- **Order status** (Pending → Processing → Shipped → Delivered)
- **Document workflow** (Draft → Review → Published)
- **Vending machine states** (NoMoney → HasMoney → Dispensing)
- **TCP connection** (Closed → Listen → Established)
- **Media player** (Playing → Paused → Stopped)

## Comparison with Other Patterns

| Pattern        | Purpose                      | Example                    |
|--------------- |------------------------------|----------------------------|
| **State**      | Change behavior based on state| Order status transitions  |
| **Strategy**   | Interchangeable algorithms   | Payment methods           |
| **Command**    | Encapsulate requests         | Button actions            |

## Characteristics

- **State Transitions**: States can transition to other states
- **Behavior Encapsulation**: Each state encapsulates its own behavior
- **Context Awareness**: States can reference the context to change state

## When to Use

✅ Object behavior depends on its state
✅ You have large conditional statements (if-else/switch) based on state
✅ State transitions are well-defined
❌ Avoid if you only have 2-3 simple states (overkill)

## Example Scenarios

- Order processing system with multiple statuses
- Document approval workflow
- Game character states (Idle, Walking, Running, Jumping)
- Network connection states