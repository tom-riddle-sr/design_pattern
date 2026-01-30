# **Iterator Pattern**

**Iterator** is a *behavioral design pattern* that provides a way to access elements of a collection sequentially without exposing its underlying representation.

## Key Concepts

- **Iterator Interface**: Defines traversal methods like `Next()` and `HasNext()`
- **Concrete Iterator**: Implements traversal logic for a specific collection
- **Aggregate/Collection Interface**: Defines a method to create an iterator
- **Client**: Uses the iterator to traverse items

## Diagram

```
Client → Iterator → Collection
           ↑            ↑
     ConcreteIterator  ConcreteCollection

Example: Playlist

Client requests iterator
        ↓
 playlist.iterator()
        ↓
 iterate songs one by one
```

## Key Advantages

- **Encapsulation**: Collection internals are hidden
- **Uniform Traversal**: Same traversal interface for different collections
- **Multiple Iterators**: Different traversal strategies can coexist
- **Single Responsibility**: Collection and traversal logic are separated

## Key Disadvantages

- **Extra Objects**: Additional iterator objects are created
- **Overhead**: Can be heavier than direct access for simple collections

## Real-World Examples

- **List/Array Iterators**: Built-in iterators in many languages
- **Database Cursors**: Iterate over query results
- **Tree Traversal**: Preorder, inorder, postorder iterators
- **File Readers**: Line-by-line iteration

## Key Differences from Other Patterns

| Pattern | Purpose | Example |
|---------|---------|---------|
| **Iterator** | Traverse collection | Playlist iterator |
| **Chain of Responsibility** | Pass requests along handlers | Support escalation |
| **Visitor** | Add operations to structures | AST processing |
| **Composite** | Treat part/whole uniformly | UI tree |

## Characteristics

- **Stable Interface**: `HasNext()` / `Next()`-style API
- **Traversal Decoupled**: Client doesn't need collection internals
- **Multiple Traversals**: Several iterators can exist simultaneously

## When to Use

✅ **Use Iterator when:**
- You need to traverse a collection without exposing internal structure
- You want multiple traversal strategies
- You want to unify traversal across different collections

❌ **Avoid when:**
- Direct access is simpler and sufficient
- Performance is extremely critical and overhead is unacceptable

## Example Scenarios

**Playlist Traversal:**
```
for it.HasNext() {
    song := it.Next()
}
```

**Tree Traversal:**
```
it := tree.InOrderIterator()
```

**File Lines:**
```
lineIt := file.LineIterator()
```