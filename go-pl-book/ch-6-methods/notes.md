**skipping the intital parts on how to create methods, things like difference between Point.Distance and (*Point).Distance**
**notes about struct embedding are in main.go file of this directory**

### Method values
- storing references to object methods in variabeles and using them later
```go

    type Point struct {X,Y int}
    func (p Point) Distance (q Point){/*....*/}

    p := Point{1,3}
    q := Point{4,6}

    distanceFromP := p.Distance     // not invoking the method, but storing the expression in some variable
    fmt.Println(distanceFromP(q))    // same as calling p.Distance(q)

```

- Another use case is situations where a package's API demands a function to be passed as an argument to another function. But the client would prefer to pass in an expression
```go

    // using time.Afterfunc() method to call the Launc() method after 10 seconds

    type Rocket struct {/*....*/}
    func (r Rocket)Launch () {/*....*/}

    r := Rocket{}
    time.Afterfunc(10 * time.Second, func () {r.Launch()})
    time.AfterFunc(19 * time.Second, r.Launch)  // this is shorter an easy to understand
```

### Method expression
- Method expression are almost simliar to method values, but in these the receiver is not predetermined. We have to pass the receiver as the first arguement to tell which object is acting as the receiver and which as the receiver
```go

    type Point struct {X,Y int}
    func (p Point) Distance (q Point){/*....*/}

    p := Point{1,3}
    q := Point{4,6}

    distance := Point.Distance      // notice it's not p.Distance, therefore we must supply the receiver explicitly

    fmt.Println(distanc(p,q))   // p will be considered as the receiver and q as the argument to the function

```

**Another expample of method expression**
```go
    type Point struct{ X, Y float64 }
    func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
    func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

    type Path []Point

    func (p Path) TranslateBy (offset Point, add bool){
        // define a general type of method expression that matches both Add and Sub
        var operation func (q Point) Point  // a function that takes in a Point and returns a Point

        if add {
            operation = Point.Add
        }else {
            operation = Point.Sub
        }
        for i := range Path {
            // based on what the 'add' boolean flag was, either Add or Sub operation will be applied to each point in the path
            path[i] = operation(path[i], offset)
        }

    }
```
