# **Factory Method**

**Factory Method** is a *creational design pattern* that defines an interface for creating objects, but lets subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

## Key Concepts

- **Creator (創建者)**: Declares the factory method that returns a product object
- **ConcreteCreator (具體創建者)**: Overrides the factory method to return a specific product
- **Product (產品)**: Defines the interface of objects the factory method creates
- **ConcreteProduct (具體產品)**: Implements the Product interface

## Diagram

```Ｖ
         +--------------+
         |    Client    |
         +------+-------+
                |
                | uses
         +------v-------+
         |   Creator    | (abstract)
         | +factoryMethod()
         | +operation()
         +--------------+
                ^
                |
                | extends
         +------+-------+
         |Concrete      |
         |Creator       |
         | +factoryMethod() -----> returns Product
         +--------------+
                                      ^
                                      |
                               +------+-------+
                               |   Product    | (interface)
                               +--------------+
                                      ^
                                      |
                               +------+-------+
                               |Concrete      |
                               |Product       |
                               +--------------+
```

## Difference from Abstract Factory

**Factory Method**: One method creates ONE product
**Abstract Factory**: One factory creates a FAMILY of related products