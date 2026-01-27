# **Proxy Pattern**

**Proxy** is a *structural design pattern* that provides a surrogate or placeholder for another object to control access to it. The proxy acts as an intermediary, adding additional functionality before or after forwarding requests to the real object.

## Key Concepts

- **Subject Interface**: Common interface shared by RealSubject and Proxy
- **RealSubject**: The actual object that performs the real work
- **Proxy**: Controls access to the RealSubject, can add extra behavior
- **Client**: Uses the proxy as if it were the real object

## Types of Proxies

1. **Virtual Proxy**: Delays expensive object creation until actually needed (lazy loading)
2. **Protection Proxy**: Controls access rights and permissions
3. **Remote Proxy**: Represents objects in different address spaces (RPC, API calls)
4. **Caching Proxy**: Stores results to avoid redundant expensive operations
5. **Logging Proxy**: Logs requests before forwarding to real object
6. **Smart Reference**: Adds additional actions when object is accessed

## Diagram

```
Client → Subject (Interface)
              ↑
              |
         _____|_____
         |         |
      Proxy    RealSubject
         |
         └─→ controls access to → RealSubject

Example: Image Loading

Client requests image
        ↓
    ImageProxy (placeholder)
        ↓
   Check if loaded?
        ↓
   No → Load from disk (expensive)
        ↓
   Yes → Return cached image
        ↓
   Display image
```

## Key Advantages

- **Access Control**: Manage who can access the real object
- **Lazy Initialization**: Create expensive objects only when needed
- **Caching**: Store results to improve performance
- **Logging & Monitoring**: Track usage without modifying real object
- **Protection**: Add security layer before accessing sensitive resources
- **Transparency**: Client uses proxy like the real object

## Key Disadvantages

- **Increased Complexity**: Extra layer of indirection
- **Potential Performance**: Slight overhead from proxy layer
- **Delayed Response**: Initial request might be slower

## Real-World Examples

- **Virtual Proxy**: Image loading (display placeholder until loaded)
- **Protection Proxy**: Access control systems, authentication
- **Remote Proxy**: gRPC, REST API clients, distributed systems
- **Caching Proxy**: Database query cache, web proxy cache
- **Smart Reference**: Reference counting, automatic resource cleanup
- **Logging Proxy**: Request/response logging, audit trails
- **Firewall Proxy**: Network security, filtering requests

## Key Differences from Other Patterns

| Pattern | Purpose | Example |
|---------|---------|---------|
| **Proxy** | Control access to object | Image lazy loading |
| **Adapter** | Convert interface | Old API → New API |
| **Decorator** | Add behavior dynamically | Coffee + Milk |
| **Facade** | Simplify interface | Unified system API |

## Characteristics

- **Same Interface**: Proxy implements same interface as RealSubject
- **Transparent**: Client doesn't know if using proxy or real object
- **Control Point**: Single point to intercept requests
- **Composition**: Proxy contains reference to RealSubject
- **Forwarding**: Eventually forwards request to real object

## When to Use

✅ **Use Proxy when:**
- Need to control access to an object
- Want lazy initialization for expensive objects
- Need to add logging, caching, or monitoring
- Require access control or authentication
- Working with remote objects or services
- Want to add reference counting or resource management

❌ **Avoid when:**
- Direct access is simpler and sufficient
- Performance overhead is unacceptable
- Object creation is not expensive

## Example Scenarios

**Virtual Proxy:**
```
// Don't load 4K video until user clicks play
videoProxy.play() → loads and plays video
```

**Protection Proxy:**
```
// Check permissions before allowing document access
documentProxy.delete() → checks auth → allows/denies
```

**Caching Proxy:**
```
// Store expensive API results
apiProxy.getData() → check cache → return or fetch
```