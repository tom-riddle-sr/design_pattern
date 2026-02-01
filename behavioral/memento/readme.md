# **Memento Pattern**

**Memento** is a behavioral design pattern that lets you capture and externalize an object’s internal state so that it can be restored later, all without violating encapsulation.

## Key Concepts

- **Originator**: The object whose state needs to be saved and restored
- **Memento**: The snapshot object that stores the Originator’s state
- **Caretaker**: Manages the mementos, but does not modify or inspect their contents

## Diagram

```
Client
       ↓
Caretaker ←→ Memento
       ↓             ↑
Originator ←────┘

Example: Text editor Undo/Redo
```

## Advantages

- **Encapsulation**: Internal state is hidden from other objects
- **Easy Undo/Redo**: Can restore to any saved state
- **Multiple Snapshots**: Supports multiple undo/redo points

## Disadvantages

- **Memory Usage**: Many snapshots can consume a lot of memory
- **Snapshot Consistency**: Complex state may require careful handling

## Real-World Examples

- **Text editor Undo/Redo**
- **Game save/load**
- **Object state snapshots**

## Comparison with Other Patterns

| Pattern        | Purpose                | Example         |
|--------------- |-----------------------|-----------------|
| **Memento**    | State snapshot/restore| Undo/Redo       |
| **Command**    | Encapsulate operation | Operation history|
| **Prototype**  | Clone object          | Copy shapes     |

## Characteristics

- **Encapsulation**: Memento’s contents are hidden from the Caretaker
- **Restorable**: Originator can restore any snapshot
- **Multiple Snapshots**: Can store many states

## When to Use

✅ Need to implement undo/redo
✅ Need to save multiple history states
❌ Avoid if state is huge or snapshots are too frequent (performance concern)

## Example Scenarios

- Text editor Undo/Redo
- Game progress save/load