# **Observer Pattern**

**Observer** is a behavioral design pattern that defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

## Key Concepts

- **Subject (Publisher)**: The object being observed, maintains a list of observers
- **Observer (Subscriber)**: Objects that want to be notified of changes
- **Attach/Detach**: Methods to add or remove observers
- **Notify**: Method to inform all observers of state changes

## Diagram

```
Subject (Publisher)
   ↓
   observers: []Observer
   ↓
   Attach(observer)
   Detach(observer)
   Notify() → calls Update() on all observers

Observer Interface
   ↑
   ├── ConcreteObserver1
   ├── ConcreteObserver2
   └── ConcreteObserver3

Flow:
1. Subject.setState(newState)
2. Subject.Notify()
3. Observer1.Update() ← notified
   Observer2.Update() ← notified
   Observer3.Update() ← notified

Example: YouTube channel (Subject) → Subscribers (Observers)
         When new video is uploaded, all subscribers get notified
```

## Advantages

- **Loose Coupling**: Subject and observers are loosely coupled
- **Dynamic Relationships**: Observers can be added/removed at runtime
- **Broadcast Communication**: One-to-many notification automatically
- **Open/Closed Principle**: Can add new observers without modifying subject

## Disadvantages

- **Unexpected Updates**: Observers don't know about each other, can cause cascading updates
- **Memory Leaks**: Forgetting to detach observers can cause memory leaks
- **Random Order**: No guaranteed order of notification
- **Performance**: Many observers can slow down notification

## Real-World Examples

- **YouTube/Newsletter subscriptions** - When new content is published, notify all subscribers
- **Event handling systems** - GUI buttons notify listeners when clicked
- **Stock market apps** - Stock price changes notify all watching investors
- **Social media notifications** - New post notifies all followers
- **MVC architecture** - Model notifies View when data changes
- **Publish-Subscribe messaging** - Message broker notifies all subscribers

## Comparison with Other Patterns

| Pattern            | Purpose                           | Communication Direction    |
|------------------- |-----------------------------------|----------------------------|
| **Observer**       | One-to-many notification          | Subject → Observers (push) |
| **Mediator**       | Centralize complex communications | Many ↔ Mediator ↔ Many    |
| **Command**        | Encapsulate requests              | Invoker → Command → Receiver |
| **Event Bus**      | Decoupled pub/sub                 | Publisher → Bus → Subscribers |

## Characteristics

- **One-to-Many**: One subject can notify many observers
- **Push Model**: Subject pushes updates to observers
- **Automatic Notification**: Observers are updated automatically when subject changes
- **Subscription Based**: Observers can subscribe/unsubscribe dynamically

## When to Use

✅ State changes in one object need to trigger updates in multiple objects
✅ Objects need to be notified without tight coupling
✅ Need dynamic add/remove of dependent objects at runtime
✅ One-to-many dependency between objects
❌ Avoid if notification overhead is too high
❌ Avoid if you need guaranteed order of updates

## Example Scenarios

- Newsletter/blog subscription system
- Real-time stock price monitoring
- Event handling in GUI frameworks
- Data binding in MVC/MVVM
- Notification systems (email, SMS, push notifications)
- Reactive programming (RxJS, React state management)