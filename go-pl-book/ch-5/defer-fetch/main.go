package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// This function fetches data from a webpage and writes the response to a file instead of Stdout.
// It derives the filename from the last component of the URL path using path.Base() function

func main() {

	fname, n, err := fetch(os.Args[1])
	if err != nil {
		fmt.Println("Errror ", err)
	}
	fmt.Fprintf(os.Stdout, "Filename = %s, Number of bytes written to file = %d\n", fname, n)
}

func fetch(url string) (filename string, n int64, errr error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending http request ", err)
		return "", 0, err
	}
	defer resp.Body.Close() // defer call

	fName := path.Base(url)
	if fName == "/" {
		fName = "index.html"
	} else {
		fName += ".html"
	}

	f, err := os.Create(fName)
	if err != nil {
		fmt.Println("Error creating a file ", err)
		return "", 0, err
	}
	// defer f.Close() 		// not recommended to close the file that we just created. In some cases, the error on file creation is not reported untill the file is closed
	// there we need the error that is possible returned while closing the file

	w, err := io.Copy(f, resp.Body)

	// since we also need the possible error occured whil copying (more important than error from f.Close() since it will tell us more about copy failure)
	// we can defer an anonymous function call which manipulates the named returns of this function
	defer func() {
		if closeErr := f.Close(); err == nil {
			errr = closeErr // err is coming from io.Copy(), we're setting the named return of errr to the closeError if io.Copy() returned no error
		}
	}()

	return fName, w, err
}

// Always remeber, since defer functions are executed after the function ends, the have the ability to manipulate the value of named returns
// this is possible due to the feature of "Closures". Functions remeber the context the were created in
