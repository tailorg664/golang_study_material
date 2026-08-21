# SOLID Principles 
This is a very important principle used for OOD and software engineering.
## Single Responsiblity Principle

- Every class should have **single responsiblity**.

## Open-Closed Principle

- Objects or entities should be **open for extension** but **closed for modification**.
- There is need of this property when there are multiple contributers and they want to extend the same function, we use this method in this way

## Liskov Substitution Principle

- Let q(x) be a property provable about objects of x of type T. Then q(y) should be provable for objects y of type S where S is a subtype of T.
- This means that every subclass or derived class should be **substitutable** for their base or parent class.

## Interface Segregation Principle

- Many client-specific interfaces are better than general-purpose interfaces.
- *Meaning* : multiple **small interfaces** are better than one large interface
- In other words, no code should be **forced** to depend on methods it doesn't use and makes the system to be decoupled and easier to refactor

## Dependency Inversion Principle

- Depends upon **Abstraction** not **Concretion**
- Use an abstraction (or interface) when declaring a method/function/etc.