# **Template Method Pattern**

**Template Method** is a behavioral design pattern that defines the skeleton of an algorithm in a base class, letting subclasses override specific steps without changing the algorithm's structure.

## Key Concepts

- **Abstract Class/Interface**: Defines the template method (algorithm skeleton)
- **Template Method**: The main method that defines the algorithm steps
- **Primitive Operations**: Abstract methods that subclasses must implement
- **Hook Methods**: Optional methods that subclasses can override

## Diagram

```
AbstractClass
   ↓
   templateMethod() {
       step1()           // Common
       step2()           // Must override
       step3()           // Common
       step4()           // Must override
   }

ConcreteClassA         ConcreteClassB
   ↓                      ↓
   step2() {...}          step2() {...}
   step4() {...}          step4() {...}

Example: Making beverages (Tea, Coffee, Hot Chocolate)
         1. Boil water (common)
         2. Brew (different for each)
         3. Pour in cup (common)
         4. Add condiments (different for each)
```

## Advantages

- **Code Reuse**: Common steps are implemented once in the base class
- **Control**: Parent class controls the algorithm structure
- **Flexibility**: Subclasses can customize specific steps without changing the overall algorithm
- **Hollywood Principle**: "Don't call us, we'll call you" - parent decides when to call subclass methods

## Disadvantages

- **Rigidity**: Subclasses must follow the algorithm structure defined by parent
- **Liskov Substitution Principle**: Can be violated if subclasses change the expected behavior
- **Limited Flexibility**: Cannot change the order of steps in the algorithm

## Real-World Examples

- **Beverage preparation** (Tea, Coffee, Hot Chocolate - same steps, different brewing/condiments)
- **Data processing** (CSV, JSON, XML - same flow: open → read → process → close)
- **Unit testing frameworks** (setUp → runTest → tearDown)
- **Game level loading** (initialize → load assets → setup entities → start)
- **Document generation** (PDF, HTML, Word - same structure: header → content → footer)

## Comparison with Other Patterns

| Pattern            | Purpose                           | Example                       |
|------------------- |-----------------------------------|-------------------------------|
| **Template Method**| Define algorithm skeleton         | Beverage preparation steps    |
| **Strategy**       | Swap entire algorithm             | Different sorting algorithms  |
| **Factory Method** | Create objects in template method | Document creation in workflow |

## Characteristics

- **Algorithm Skeleton**: The template method defines the fixed sequence of steps
- **Inversion of Control**: Parent class calls subclass methods (Hollywood Principle)
- **Code Reuse**: Common steps are implemented once in the base class
- **Customization Points**: Subclasses override specific steps to customize behavior

## When to Use

✅ Multiple classes have similar algorithms with minor differences
✅ You want to control the algorithm structure while allowing customization
✅ Common behavior can be factored out into a base class
✅ You want to avoid code duplication across similar classes
❌ Avoid if subclasses need to change the algorithm order or structure

## Example Scenarios

- Building different types of reports (PDF, HTML, Excel) with same structure
- Processing different file formats (CSV, JSON, XML) with same workflow
- Implementing test frameworks with consistent setUp/test/tearDown flow
- Creating different types of beverages with same preparation steps