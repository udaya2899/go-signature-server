## Go Challenge
### Three Underrated Features in Go

1. **Documentation with comments**:
Commenting above functions and variables, which is warned by `vet` for exported functions and variables, generates documentation. Amazingly simple. This is so simple compared to other language documentation generation ways. Writing a short concise comment of every function is a must in my opinion.

```go
// SomeFunc is a function that does something. Note that when not giving anything, it'll not do anything.
func SomeFunc() {
    return
}
```

2. **init()**:
Statements inside an init() function is executed when the package is called. Every package's initial setup is best done in init() and understanding of the lifecycle is important to write idiomatic Go code. A simple graphical example below that helped me learn what init() is:

![init illustration](https://astaxie.gitbooks.io/build-web-application-with-golang/en/images/2.3.init.png?raw=true)
Picture Credits: StackOverflow

Sample Usage: \
```go
package foo

var ImportantVar string

func init() {
    ImportantVar, err = bar()
    if err != nil {
        log.Panicf("Cannot get important var")
    }
}

func bar() (string, error) {
    // get string from somewhere
    return str
}
```

`ImportantVar` can now be used directly by doing `foo.ImportantVar`. Note that errors inside init can be handled only by panic or exiting



1. **Using Object Oriented-ness with structs**:
Go is not generally listed amongst the list of OOP languages because it doesn't directly implement classes. Go implements (almost) the same without classes and it's quite underrated, especially from people who switch from non-OOP languages like JavaScript and Python. 

A simple OOP implementation can be implemented like (a trivial example below):
```go
package service

type Person struct {
    Name string
    Age string
}

// Creates a new instance of Person 
func New(name, age string) *Person {
    return &Person{
        Name: name,
        Age: age
    }
}

// returns the length of name of the instantiated object - the function can only be called by the object that was created with New()
func (p Person) GetNameLength() string {
    return len(p.Name)
}
```

### When Go shoots me in my foot?

**Default arguments** - Go does not provide an option to provide default values to arguments when not provided explicitly. Although this improves readability, a simple way to provide a default value would've opened the door to more possibilities

### Usage of Unsafe

> NOTE: I haven't gotten the need to use Unsafe directly in the code I've written, yet, so my knowledge about this package is more on what I've read from multiple sources, the documentation and from devs.

To my knowledge, the best use of Unsafe from a developer's point of view is *to convert a struct* of type A to type B if A and B both have same structure underlying inside.

Using `unsafe.Pointer` \

For eg:
```go
type A struct {
    numberA int
}

type B struct {
    numberB int
}

func foo() {
    a := A{
        numberA: 2
    }
    
    b := *(*B)(unsafe.Pointer(&a))

    // b is now of type B
}
```
`unsafe.Offsetof` can be used to find the offset of a field in a struct. By doing this, the value of the field can be directly altered in memory without allocating more.

`unsafe.Sizeof` is used to iterate over an array to get the size of element to iterate over

To summarize, unsafe package directly hits on one of the best features of Go i.e its type safety, smart memory management etc. So, it should be used by someone who knows what it does exactly (experts) and still, things might go wrong. unsafe is to be used only when the necessity is met.