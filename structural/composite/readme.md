# **Composite Pattern**

**Composite** is a *structural design pattern* that allows you to compose objects into tree structures to represent part-whole hierarchies. It lets clients treat individual objects and compositions of objects uniformly.

## Key Concepts

- **Component**: Common interface for both leaf and composite objects
- **Leaf**: Represents end objects (files in our example)
- **Composite**: Contains other components (files and directories)
- **Operation**: Methods called uniformly on all components

## Diagram

```
        FileSystemComponent (Interface)
             /           \
          File          Directory
        (Leaf)         (Composite)
                        /    |    \
                     File  File  Directory
                              (Tree Structure)
```

## Key Advantages

- Represents hierarchical structures (tree structures)
- Allows treating individual and composite objects uniformly
- Simplifies client code (no need to check types)
- Easy to add new component types
- Promotes single responsibility principle
- Makes client code simpler and more elegant

## Real-World Examples

- File systems (files and directories)
- GUI components (containers and widgets)
- Organization hierarchies (departments and employees)
- Document structures (chapters and sections)
- Graphics (groups of shapes)
- Menu systems (menus and menu items)

## Characteristics

- **Recursive composition**: Components can contain other components
- **Uniform treatment**: Single interface for diverse object types
- **Transparency**: Clients work with interface, not concrete classes
- **Flexibility**: Easy to add new component types

- Graphics rendering (shapes × graphics engines)