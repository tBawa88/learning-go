## Pointers 
- A pointer stores the address of a variable, where that variable is stored inside memory. Not every value has address but every variable does.
- Zero value of a pointer is 'nil'
- Functions can return address to the locally created variables
- Pointers are used to read or update the value of a variable indirectly

### echo4.go
- This program is demonstrating the use of "flags" package which works with pointers and returns pointers to the created values


- "flag" package is used to create flags that can be passed as command line flags while executing the .exe file
- flag.String() is creating a flag named -O, it's default value is ",", followed by it's description. It then returns a *string, which points to the value of that flag
- flag.Bool() is creating boolean flag, which by default is flase, but if included into the command line arguments will be set to true. Returns a *bool which points to the value of that flag
- flag.Parse() to use the values of flags, we must include this at the top of the programm.
- `cla := strings.Join(flag.Args(), *sep)` This is joining all the command line arguments into a single string using the value of the -O flag
- `if !*n {
		fmt.Println()
	}` 
- if -no flag is included as the flag, then no new line will be added, new line will be added by default  

## new() function
- This is another way of creating pointers of a certain types
`p := new(int)`
- Here 'p' is the pointer to an unnamed variable of type int, by default initialized to zero value
- `*p = 40`

```go
    func newInt () *int {
        return new(int)
    }

    func newInt () *int {
        var dummy int
        return &dummy
    }
```
- Both of these functions are doing exact same thing. Returning a pointer to a local variable of type int
**NOTE** : new is a function and not a keyword, therefore it can be redefined and used as function arguments 


## Type declarations
- There are situations where variables of same type are representing completely different things in a program. Now it could lead into accidental mixing of those values since they're of same type
- To solve this potential problems, and improve code structure, Go introduced **type** declarations. 
- It is used to create a new **named** type which extends a base type (int, string, float64, bool) as it's underlying type. 
```go
    type Celcius float64
    type Fahrenheit float64
```
- Both of these types  have same underlying type but they're now a completely differnt type. Meaning you cannot compar their variables
- Another benefit of creating named types like this is that you can define receiver functions, which creates additional functionality for the base type (define methods that are only available to variables of a certain type)


## Conversions of types 
- A conversion from one type to another is allowed if both of them are extending as long as they're both extending the same base type
- For every type T, there is a corresponding conversion operation T(x) which converts the value x to type T 


## exercise 2.1 
- Create a new package tempconv, which has Kelvin and Celcius scale, where 0 Kelvin is -273.15 Celcius. 
- **SOLUTION** : to create a package in some folder and use it inside the project first me first initialize the module 
- `go mod init example.com/go-course` -> this creates a go.mod file in the root of the folder which makes importing the modules using their import paths easier
- `import "tbawa/go-course/go-pl-book/ch-2-programStructure/tempconv"` like this you can provide that path to the package be prefixing it with **tbawa/go-course/pathToPackageFiles**