# **Bridge Pattern**

**Bridge** is a *structural design pattern* that separates an abstraction from its implementation, allowing them to vary independently. It's like building a bridge between two independent dimensions of variation.

## Key Concepts

- **Abstraction**: High-level logic that uses the implementation (e.g., Note types)
- **Implementation**: Low-level concrete details (e.g., Storage methods)
- **Bridge**: The reference to implementation held by abstraction
- **ConcreteAbstraction**: Specific abstraction types (e.g., SimpleNote, EncryptedNote)
- **ConcreteImplementation**: Specific implementation types (e.g., LocalStorage, CloudStorage)

## Diagram

```Abstraction Layer
   (What to do)
   +---------+---------+
   |         |         |
Simple    Encrypted  Regular
Note       Note      Note
   |         |         |
   +----+----+----+----+
        | Bridge
        v (holds reference)
   +----------+
   | Storage  |
   |(How to)  |
   +----------+
   |    |     |
Local Cloud  Memory
Store Store  Store
         +--------+
```
## Key Advantages

- Allows reuse of existing classes with incompatible interfaces
- Promotes single responsibility principle
- Flexible and non-invasive (don't modify original classes)
- Evoids class explosion (N × M combinations → N + M classes)
- Decouples abstraction from implementation
- Promotes single responsibility principle
- Easy to extend both dimensions independently
- More flexible than deep inheritance hierarchie
- Legacy system integration
- Third-party library integration
- GUI frameworks (shapes × colors, controls × platforms)
- Cross-platform applications (UI × OS)
- Database systems (query types × storage engines)
- Note/document systems (note types × storage methods)
- Remote controls × devices
- Graphics rendering (shapes × graphics engines)