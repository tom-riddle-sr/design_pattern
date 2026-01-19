# **Abstract Factory**

**Abstract Factory** is a *creational design pattern* that defines a factory interface for creating families of related objects without specifying their concrete classes. Clients depend only on the factory interface, so different concrete factories can supply matching product variants while keeping client code unchanged.

## Diagram

```
	   +--------------+
	   |    Client    |
	   +------+-------+
			|
			| calls
	   +------+-------+
	   | Abstract    |
	   |  Factory    |
	   +------+------+ 
			| creates
	 +--------+---------+
	 |                  |
	 v                  v
 +----+-----------+  +---+-----------+
 | ConcreteFactoryA|  |ConcreteFactoryB|
 +----+-----------+  +---+-----------+
	 | returns            | returns
	 v                    v
 +----+-----+        +----+-----+
 | ProductA1|        | ProductB1|
 +----------+        +----------+
	 |                    |
	 | returns            | returns
	 v                    v
 +----+-----+        +----+-----+
 | ProductA2|        | ProductB2|
 +----------+        +----------+
	 ^                    ^
	 | used via interface |
	 +--------------------+
```