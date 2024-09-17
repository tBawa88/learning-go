
**Interface values**
- Interfaces are not just types that define method contracts for other types to satisfy
- Just like other types, we can define variables of interface types (even though they don't store property values on em)
- An interface type has 2 components 1. type 2. value. For a zero value interface, both of these components are set to nil

```go
    var w io.Writer     // io.Writer is an interface, and declared a variabel named w with Writer interface as it's type
    w = os.Stdout
    w = new (bytes.Buffer)
    w = nil
    // all 3 values are valid values for w


    w = os.Stdout
    // after this type of w is set to *os.File , this is possible since *os.File satisfies io.Writer interface

    w.Write([]byte("Hello there"))
    // this calls (*os.File).Write() method and prints "Hello there" in the terminal
```
