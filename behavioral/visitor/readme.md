# **Visitor Pattern**

**Visitor** is a behavioral design pattern that allows you to separate algorithms from the objects on which they operate. It lets you add further operations to objects without modifying their structure.

## Key Concepts

- **Visitor Interface**: Declares a visit method for each type of concrete element.
- **Concrete Visitor**: Implements the operations defined in the visitor interface.
- **Element Interface**: Declares an accept method that takes a visitor object.
- **Concrete Element**: Implements the accept method to call the visitor's method corresponding to its class.
- **Object Structure**: A collection of elements that can be iterated and visited by visitors.

## Diagram

```
Visitor Interface
   ↓
   visitConcreteElementA()
   visitConcreteElementB()

ConcreteVisitorA         ConcreteVisitorB
   ↓                          ↓
   visitConcreteElementA()    visitConcreteElementA()
   visitConcreteElementB()    visitConcreteElementB()

Element Interface
   ↓
   accept(visitor: Visitor)

ConcreteElementA         ConcreteElementB
   ↓                          ↓
   accept(visitor: Visitor)   accept(visitor: Visitor)

ObjectStructure
   ↓
   elements: []Element
   accept(visitor: Visitor)
```

## Advantages

- **Open/Closed Principle**: Add new operations without modifying existing classes.
- **Single Responsibility Principle**: Separate algorithms from the objects they operate on.
- **Flexibility**: Add new visitors to extend functionality.

## Disadvantages

- **Complexity**: Adding new element types requires changes to all visitors.
- **Dependency**: Tight coupling between visitors and elements.

## Real-World Examples

- **Compilers**: Abstract syntax tree traversal with different operations (e.g., type checking, optimization).
- **Serialization**: Converting objects into different formats (e.g., JSON, XML).
- **UI Component Rendering**: Rendering different UI components with specific behaviors.
- **File System Operations**: Performing operations like searching, compressing, or encrypting files.

## Comparison with Other Patterns

| Pattern            | Purpose                           | Key Feature                  |
|------------------- |-----------------------------------|------------------------------|
| **Visitor**        | Add operations to objects         | Separate algorithms from data|
| **Strategy**       | Swap entire algorithm             | Encapsulates algorithms      |
| **Command**        | Encapsulate requests              | Decouples sender and receiver|

## Characteristics

- **Double Dispatch**: The accept method allows the visitor to determine the concrete element type at runtime.
- **Separation of Concerns**: Operations are separated from the objects they operate on.
- **Extensibility**: New operations can be added by creating new visitors.

## When to Use

✅ Need to perform multiple unrelated operations on an object structure.
✅ Want to add new operations without modifying existing classes.
✅ Need to centralize operations on a complex object structure.
❌ Avoid if the object structure is unstable or frequently changes.

## Example Scenarios

- Traversing and processing a file system (e.g., searching, compressing, encrypting files).
- Implementing operations on a complex object model (e.g., compilers, interpreters).
- Rendering UI components with different behaviors.
- Exporting objects to various formats (e.g., JSON, XML, CSV).