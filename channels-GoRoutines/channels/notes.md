The goal is to create a program that pings multiple sites by sending them an HTTP get request to check whether these sites are up or not

## First iteration 
```go
    for _, url : range urls {
        checkStatus(url)    //the blocking call
    }
```
- When we're making the request like this in a loop, our program waits untill the response from that url comes back before moving on the next iteration.
- This approach is quite inefficient, imagine if had thousands of links to check, the program will take a while to check the status of each link
- Instead of this Serial approach of checking each link one by one, it'd be great if you send a request to all the links at the same time, and whichever request resolves, we could print it's result in the terminal

## Go Routines
- When we create a normal go program, by default a single go routine get's created inside it. This routine executes the code in a top down fashion
- When it gets to the http.Get() it has to wait there and can't do anything else. We refer to this as a "blocking call"
- We can launch a new go routine by placing the **keyword "go"** in front of the blocking call
```go 
    for _, url := range urls {
        go checkStatus(url)     //launched a new go routine that will execute this function
    }
```
- When the main routine hits this go keyword, it launches another routine who's whole job is to execute this function
- When this newly created go routine hits `res, err := http.Get(url)`, the blocking call, it goes to sleep 
- Before goin to sleep, it also emitts en event which notifies any other routines that might be waiting for it's execution to go ahead and execute their own job
- The  main routine then goes ahead and executes it's next iteration of for loop 
- Then in the next iteration, again another go routine is launched which repeats the same thing. Hits the blocking call, goes to sleep, notifies other routines to go ahead 
- 