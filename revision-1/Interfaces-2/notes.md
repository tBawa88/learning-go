### Goal
- TO better understand how interfaces are implemented implicitly in go by buiding a simple program that sends an HTTP request and prints the response to the terminal


### net/http package
```go
    resp, err := http.Get("www.something.com")
    // resp has a body property which is of type io.ReadCloser interface
    // THis ReadClose interface further implements 2 more interface types "Reader" "Closer"
```

**Reader Interface**
- Defines a contract that a type must have a function called `Read(p []byte) (n int, err error)`
- This means that the resp.body must have a method called resp.body.Read(p [] byte)
- It reads all the bytes into this byte slice called p
- We can try converting this byte slice into a string and see what kind of response we got

**Closer Interface**
- Defines a contract that a type must have a function called `Close() error`


### Solution 1
- Create a byte slice of some size and pass it to the Read() function of Body object
- All the data from the body will be read into this byte slice which we can convert into a string
```go
    bSlice := make([]byte, 99999)

    res, _ := http.Get("http://localhost:3000/home")

    n, _ := res.Body.Read(bSlice)

    fmt.Println(string(bSlice))
```
- BUT there is a better way instead of creating a byte slice of 9999 size every time, we can use **io.Copy** function from **io** package

### Solution 2
- `io.Copy(target Writer, src Reader)`
- This function writes all the data from the src to target
- 'src' is of type Reader interface, meaning it must be some object that has `Read(b []byte) (int, err)` receiver function
- 'target' is of type Writer inteface, meaning it must be some object that has `Write(b []byte) (int, err)` jreciever function

- Now we must find 2 such objects
1. Reader will come from res.Body
2. Writer will come from os.Stdout (it has Write()) on it which writes the byte slice onto the terminal output

```go
func writeToTerminal(res *http.Response) {
	io.Copy(os.Stdout, res.Body)
	fmt.Println("")
}
```

### Solution 3 - Write the response data to a file
- Create a new struct type call it **fileWriter**
```go
    type fileWriter struct {
        filename string
    }

    func (f fileWriter) Write(b []byte) (int, error){
        os.Remove(f.filename)
        err := os.WriteFile(f.filename, b, 0666)
        return len(b), err
    }
```
- This type is not implementing the Writer interface since it has a Write() receiver function
- This Write() function will automatically be called by the io.Copy() function and we'll have access to the pre-populated byte slice
- Writing that byte slice to a file is easy using the "os" package


## Assignment
- Write a program that takes a filename as an input from the terminal
- Opens that file
- Reads it
- And writes the contents of that file to a new Copy of that file


```go
    type fileWriter struct {}

    func (f fileWriter) Write(b []byte) (int, error){
        fmt.Println(string(b))
    }

    func main () {
        fileName := os.Args[1]
        file := os.Open(filename)
        var f fileWriter
        io.Copy(f, file)
    }
```

