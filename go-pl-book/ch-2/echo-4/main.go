// Echos the command line arguements but also watches for a couple of flags passed
// -n flag means ommit the trailing newline or not
// -s flag looks whether a seperator has been provided or not

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //defualt value is given as false, but user can overwrite it by passing it
var s = flag.String("s", " ", "seperator")             //default value is a "space" but user can provide a custom seperator to join the arguments

func main() {
	flag.Parse()                                    //Parses all the flags passed from the os.Args[1:], must be used after all the flags have been defined
	fmt.Printf("%s", strings.Join(flag.Args(), *s)) //flag.Args() returns the non-flag values of the arguments passed to the program

	if !*n {
		fmt.Println()
	}
}
