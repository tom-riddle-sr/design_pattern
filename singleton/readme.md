# **Singleton**

**Singleton** is a *creational design pattern* that ensures a class has only one instance and provides a global point of access to it.

## Key Concepts

- **Singleton (單例)**: Ensures only one instance exists
- **Private Constructor**: Prevents external instantiation
- **Static Instance**: Holds the single instance
- **Global Access Point**: Provides access to the instance

## Diagram

```
         +--------------+
         |  Singleton   |
         +--------------+
         | -instance    | (static)
         +--------------+
         | +GetInstance()| (static)
         | -Singleton() | (private constructor)
         +--------------+
```
◊
## Key Advantages

- Controlled access to sole instance
- Reduced namespace pollution
- Permits refinement of operations and representation
- Lazy initialization possible

## Use Cases

- Database connections
- Configuration managers
- Logger systems
- Thread pools
- Cache managers
