# **Adapter Pattern**

**Adapter** is a *structural design pattern* that allows objects with incompatible interfaces to collaborate. It works like a real-world adapter (e.g., USB converter, power adapter).

## Key Concepts

- **Target Interface**: The interface the client expects
- **Adaptee**: The existing incompatible interface/class
- **Adapter**: The class that bridges Target and Adaptee
- **Client**: Uses the Target interface through the Adapter

## Diagram

```
         +-----------+
         |  Client   |
         +-----------+
              | uses
              v
      +---------------+
      | Target Iface  |
      +---------------+
              ^
              | implements
         +----------+
         | Adapter  |
         +----------+
         | adaptee  |
         +----------+
              | uses
              v
         +--------+
         | Adaptee|
         +--------+
```
## Key Advantages

- Allows reuse of existing classes with incompatible interfaces
- Promotes single responsibility principle
- Flexible and non-invasive (don't modify original classes)
- Easy to swap different adapters

## Use Cases

- Legacy system integration
- Third-party library integration
- Format conversion (JSON ↔ XML)
- Different API versions compatibility
- Hardware device drivers
