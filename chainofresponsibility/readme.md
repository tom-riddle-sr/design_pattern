# **Chain of Responsibility Pattern**

**Chain of Responsibility** is a *behavioral design pattern* that lets you pass requests along a chain of handlers. Each handler decides either to process the request or to pass it to the next handler in the chain.

## Key Concepts

- **Handler Interface**: Defines a method to handle requests and a reference to the next handler
- **Concrete Handlers**: Implement specific handling logic and decide whether to pass the request onward
- **Client**: Sends the request to the first handler in the chain

## Diagram

```
Client → HandlerA → HandlerB → HandlerC → ...
          |           |           |
       handle      handle      handle
       or pass     or pass     or pass

Example: Support Ticket

Client submits ticket
        ↓
  Level1Support (simple issues?)
        ↓ no
  Level2Support (technical issues?)
        ↓ no
  Manager (policy/exception?)
        ↓
      resolved
```

## Key Advantages

- **Loose Coupling**: Client does not need to know which handler will process the request
- **Flexible Chain**: Add, remove, or reorder handlers at runtime
- **Single Responsibility**: Each handler focuses on one responsibility
- **Extensible**: New handlers can be added without changing existing code

## Key Disadvantages

- **Unhandled Requests**: A request might go through the entire chain without being handled
- **Debugging Complexity**: Tracking the flow can be harder
- **Potential Performance**: Long chains can add overhead

## Real-World Examples

- **UI Event Bubbling**: Events propagate through UI components
- **Logging Pipelines**: Debug → Info → Warn → Error handlers
- **Middleware**: Web server request middleware chain
- **Authorization**: Role checks in sequence
- **Validation**: Input validators in order

## Key Differences from Other Patterns

| Pattern | Purpose | Example |
|---------|---------|---------|
| **Chain of Responsibility** | Pass requests through handlers | Support ticket escalation |
| **Observer** | Notify multiple listeners | Event subscription |
| **Strategy** | Swap algorithm behavior | Sorting strategy |
| **Decorator** | Add behavior dynamically | Coffee + Milk |

## Characteristics

- **Handler Chain**: Each handler knows only the next handler
- **Optional Handling**: Handlers may handle or forward
- **Runtime Configuration**: Chain can be built dynamically
- **Decoupled Client**: Client depends only on handler interface

## When to Use

✅ **Use Chain of Responsibility when:**
- Multiple objects can handle a request
- You want to avoid coupling sender to receiver
- You want flexible, runtime-configurable processing
- You need a pipeline of checks or transformations

❌ **Avoid when:**
- Only one handler ever processes the request
- Requests must always be handled with strict ordering guarantees
- Chain length could be very large with tight performance constraints

## Example Scenarios

**Validation Pipeline:**
```
// Validate input step-by-step
validatorChain.handle(input)
```

**Logging Levels:**
```
// Only logs at or above configured level
loggerChain.handle(logMessage)
```

**Authorization:**
```
// Role checks from least to most privileged
authChain.handle(request)
```