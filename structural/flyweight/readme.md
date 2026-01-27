# **Flyweight Pattern**

**Flyweight** is a *structural design pattern* that uses sharing technique to support large numbers of fine-grained objects efficiently. It reduces memory usage by sharing common state between multiple objects instead of storing all data in each object.

## Key Concepts

- **Intrinsic State**: Shared, immutable data that can be stored in the Flyweight (e.g., tree type, character font)
- **Extrinsic State**: Context-dependent data that varies between objects (e.g., position, size)
- **Flyweight Factory**: Manages and shares Flyweight objects, ensuring objects are created only once
- **Shared Pool**: Caches created objects to avoid duplicate instantiation

## Diagram

```
Client → FlyweightFactory → Flyweight (Intrinsic State)
              ↓                    ↑
         [Cache Pool]      Pass Extrinsic State
         
Example: Forest Rendering

Without Flyweight:
  1,000,000 trees = 1,000,000 separate objects
  Memory: ~1GB

With Flyweight:
  1,000,000 trees = 3 shared tree types + 1,000,000 coordinates
  Memory: ~10MB
  
TreeType (Intrinsic - Shared):
  - name: "Oak"
  - color: "Green"
  - texture: [image data]

Tree (Extrinsic - Individual):
  - x: 100
  - y: 200
  - treeType: → shared TreeType
```

## Key Advantages

- **Massive Memory Reduction**: Share common data across many objects
- **Performance Improvement**: Reduced object creation overhead
- **Centralized Management**: Factory pattern manages shared objects
- **Scalability**: Can handle millions of objects efficiently

## Key Disadvantages

- **Increased Complexity**: Need to separate intrinsic and extrinsic state
- **Runtime Cost**: Additional logic to manage external state
- **Thread Safety**: Shared objects require careful synchronization

## Real-World Examples

- Text editors (share character objects, vary position)
- Game engines (trees, bullets, particles with thousands of instances)
- Graphics systems (icons, sprites, textures)
- String pools (Java String interning)
- Connection pools, thread pools
- Browser DOM rendering (shared styles)
- Font rendering systems

## Key Differences from Other Patterns

| Pattern | Purpose | Example |
|---------|---------|---------|
| **Flyweight** | Share objects to save memory | Share font objects |
| **Singleton** | Ensure single instance | Global config |
| **Object Pool** | Reuse expensive objects | Connection pool |
| **Prototype** | Clone objects | Deep copy objects |

## Characteristics

- **Large Quantity**: Best for scenarios with massive similar objects
- **State Separation**: Intrinsic state shared, extrinsic state independent
- **Immutability**: Intrinsic state is typically immutable
- **Factory Management**: Factory ensures object sharing
- **Memory vs CPU Trade-off**: Save memory at the cost of computation

## When to Use

✅ **Use Flyweight when:**
- Application uses a huge number of objects
- Storage costs are high due to quantity
- Most object state can be made extrinsic
- Many groups of objects can be replaced by few shared objects

❌ **Avoid when:**
- Few objects needed
- Objects don't share common state
- Extrinsic state is complex to manage