## Object oriented Programming


```go
✅ 1. Encapsulation
Encapsulation means hiding internal details and exposing only what's necessary.

Go achieves it using:
Capitalized names for public (exported) identifiers.

Lowercase names for private (unexported) identifiers.

Example:

package user

type User struct {
    Name  string  // Exported
    email string  // Unexported
}

func (u *User) GetEmail() string {
    return u.email
}

func (u *User) setEmail(e string) {
    u.email = e
}

✅ 2. Abstraction
Abstraction means exposing only essential features and hiding the complexity.

Go achieves it using:
Interfaces: Define behavior without exposing implementation.


type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof"
}

func MakeItSpeak(a Animal) {
    fmt.Println(a.Speak())
}
MakeItSpeak() doesn't care which animal — just that it can speak.

✅ 3. Polymorphism
Polymorphism allows different types to be treated as the same interface.

Go achieves it using:
Interface implementation without explicit declaration.


type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
Both Circle and Square can be passed to PrintArea().

✅ 4. Inheritance (via Composition)
Go does not support classical inheritance, but it uses composition, which is more flexible.

Composition example:

type Person struct {
    Name string
}

func (p Person) Greet() {
    fmt.Println("Hello,", p.Name)
}

type Employee struct {
    Person    // Embedded type (composition)
    EmployeeID string
}
Now Employee inherits Greet() method from Person.


e := Employee{
    Person: Person{Name: "Rishabh"},
    EmployeeID: "123",
}
e.Greet() // Hello, Rishabh

```