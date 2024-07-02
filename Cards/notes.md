## OOP approach vs Go approach
- In traditional oop langs, when we want to create our own data type, we either create a class or a struct or something along those lines. 
- Then we define instance methods which are acessible to all the instances of that class/struct
- But in Go there is no concept of a class. To define our own data type we extend a base type using the **type** keyword

`type deck [] string`
This defines a new type called **deck** which is an array of string type. 

## Defining methods for type deck (receiver functions)
```go
    func (d deck) printDeck() {
        for _, card := range d{
            fmt.Println(card)
        }
    }
```
- This is called a receiver function. Before the name of the function, we define the type to which this func will be related to 
- It receives a copy of the object that calls it 
  `func (t type ) functionName() <returnType> {}`

## Calling methods on type variables
```go
    myDeck := deck{"Jack of Spades", newCard(), "Queen of Hearts"}

    myDeck.printDeck(); //myDeck will be passed as a copy to this funciton, therfore any changes made will not reflect in the original
```
**Pointer receiver** :
```go
// This is called pointer receiver. Any object that calls this method will have it's address passed to this and stored inside the variable d
func (d *deck) addNewCard(card string) {
	*d = append(*d, card)
}
```

