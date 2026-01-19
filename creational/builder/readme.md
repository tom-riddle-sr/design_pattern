# **Builder**

**Builder** is a *creational design pattern* that separates the construction 
of a complex object from its representation, allowing the same construction 
process to create different representations. It uses a Director to define 
the construction steps and sequence.
## Diagram

```
         +--------------+
         |    Client    |
         +------+-------+
                |
                | uses
         +------v-------+
         |   Director   |
         +------+-------+
                |
                | orchestrates
         +------v-------+
         |   Builder    | (interface)
         +--------------+
                ^
                |
                | implements
         +------+-------+
         |Concrete      |
         |Builder       |
         +------+-------+
                |
                | builds
         +------v-------+
         |   Product    |
         +--------------+
         
         

```