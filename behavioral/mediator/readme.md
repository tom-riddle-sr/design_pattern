# **Mediator Pattern**

**Mediator** is a behavioral design pattern that reduces chaotic dependencies between objects by restricting direct communications and forcing them to collaborate only via a mediator object.

## Key Concepts

- **Mediator**: Defines the interface for communication between colleague objects
- **Concrete Mediator**: Implements the mediator interface and coordinates colleagues
- **Colleague**: Objects that communicate through the mediator instead of directly
- **Decoupling**: Colleagues don't know about each other, only the mediator

## Diagram

```
         Mediator
            ↑
    ┌───────┼───────┐
    ↓       ↓       ↓
Colleague1 Colleague2 Colleague3

Without Mediator (Chaos):
A ←→ B
A ←→ C
A ←→ D
B ←→ C
B ←→ D
C ←→ D
(N objects = N*(N-1)/2 connections)

With Mediator (Clean):
A → Mediator ← B
    ↑     ↓
    C ← → D
(N objects = N connections)

Example: Chat room mediator between users
```

## Advantages

- **Decoupling**: Components don't need to know about each other
- **Centralized Control**: All communication logic is in one place
- **Simplified Maintenance**: Easier to modify communication logic
- **Reduced Dependencies**: From N² to N connections
- **Single Responsibility**: Each component focuses on its own logic

## Disadvantages

- **God Object**: Mediator can become too complex and hard to maintain
- **Single Point of Failure**: If mediator fails, all communication breaks
- **Performance**: All communication goes through one object (potential bottleneck)

## Real-World Examples

- **Chat Room** - Users send messages through the chat room, not directly
- **Air Traffic Control** - Planes communicate through ATC, not with each other
- **GUI Dialog** - Form components interact through dialog controller
- **Smart Home Hub** - Devices communicate through central hub
- **MVC Controller** - Mediates between Model and View

## Comparison with Other Patterns

| Pattern        | Purpose                        | Communication          |
|--------------- |--------------------------------|------------------------|
| **Mediator**   | Centralize complex communications | Many-to-many via mediator |
| **Observer**   | Notify multiple objects        | One-to-many direct     |
| **Facade**     | Simplify subsystem interface   | One-way only           |
| **Chain of Responsibility** | Pass request along chain | Sequential           |

## Characteristics

- **Centralized Communication**: All interactions go through the mediator
- **Loose Coupling**: Components don't reference each other directly
- **Complexity Reduction**: From N² to N relationships
- **Coordinator Role**: Mediator knows about all colleagues and orchestrates their collaboration

## When to Use

✅ Components have complex, tangled dependencies
✅ Difficult to reuse components because they're too coupled
✅ Need to centralize complex communication logic
✅ Many-to-many relationships between objects
❌ Avoid if communication is simple (adds unnecessary complexity)
❌ Avoid if only one-to-many relationships (use Observer instead)

## Example Scenarios

- Chat application with multiple users
- Air traffic control system
- GUI forms with interdependent controls
- Smart home automation system
- Game event system