- **Panics** are runtime errors in go. They could occur naturally for eg: an array out of bounds access, or a nil pointer dereference while walking a linked list etc.
- Think of Panics as uncaught errors in JS
- We also have to ability to directly called the built-in panic function to generate a panic.
- A panic is often best thing to do when an impossible situtation occurs

```go
// btw switch satements are just like if statements, they allow us to do shorthand assginments as the first statement, and use that value as the condition in the 2nd statement
    switch s := suit(drawCard()); s {
        case "Spades":
        case "Diamonds":
        case "Clubs":
        case "Hearts":
        default:
            panic(fmt.Sprintf("invalid suit %q", s))
    }
```
