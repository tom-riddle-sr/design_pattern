# **Prototype**

**Prototype** is a *creational design pattern* that lets you copy existing objects without making your code dependent on their classes.

## Key Concepts

- **Prototype (原型)**: Declares an interface for cloning itself
- **ConcretePrototype (具體原型)**: Implements the cloning interface
- **Client (客戶端)**: Creates a new object by asking a prototype to clone itself

## Diagram

```
         +--------------+
         |    Client    |
         +------+-------+
                |
                | clone()
         +------v-------+
         |  Prototype   | (interface)
         | +Clone()     |
         +--------------+
                ^
                |
                | implements
         +------+-------+
         |Concrete      |
         |Prototype     |
         | +Clone()     |
         +--------------+
```

## Key Advantages

- Avoid expensive creation operations
- Create objects without knowing their exact classes
- Reduce subclassing overhead
