# **Decorator Pattern**

**Decorator** is a *structural design pattern* that allows you to attach new responsibilities to an object dynamically, providing a flexible alternative to subclassing. It lets you add behavior to objects without modifying their underlying code.

## Key Concepts

- **Component**: Interface that defines the object we want to add behavior to
- **Concrete Component**: The original object (e.g., SimpleCoffee)
- **Decorator**: Abstract class that implements the Component interface and wraps a Component
- **Concrete Decorator**: Adds specific behavior/responsibility to the wrapped component

## Diagram

```
     Component (Interface)
       /             \
    Coffee      CoffeeDecorator
                      |
         ______________|______________
         |      |      |      |
      Milk  Sugar  Whip  Chocolate
   (Decorator) (Decorator) (Decorator) (Decorator)

Simple Coffee (¥100)
     ↓ wrap with Milk
Coffee + Milk (¥150)
     ↓ wrap with Sugar
Coffee + Milk + Sugar (¥160)
     ↓ wrap with Whip
Coffee + Milk + Sugar + Whip (¥180)
```

## Key Advantages

- **No Subclass Explosion**: Don't need CoffeeWithMilk, CoffeeWithMilkAndSugar, etc.
- **Flexible Combination**: Dynamically add multiple features in any order
- **Single Responsibility**: Each Decorator handles one specific behavior
- **Open/Closed Principle**: Open for extension, closed for modification
- **Dynamic Behavior**: Add features at runtime, not compile time
- **Composable**: Stack multiple decorators together

## Real-World Examples

- Coffee shops (coffee + toppings)
- GUI components (scroll bars, borders, shadows)
- Text editors (text + bold, italic, underline)
- Stream I/O (file → compress → encrypt)
- Logging systems (add logging to any method)
- Caching layers (data access with cache)
- UI components (buttons with icons, badges, etc.)

## Key Differences from Other Patterns

| Pattern | Purpose | Example |
|---------|---------|---------|
| **Decorator** | Add new responsibility | Coffee + Milk |
| **Adapter** | Convert interface | Old API → New API |
| **Composite** | Tree structure | Files + Folders |

## Characteristics

- **Wrapping**: Each decorator wraps a component
- **Same Interface**: Decorators implement the same interface as the component
- **Stacking**: Multiple decorators can be stacked
- **Transparent to Client**: Client doesn't know about decorators
- **Recursive Composition**: Decorators can wrap other decorators