package main

import (
	"fmt"

	"tbawa/go-course/go-pl-book/ch-2-programStructure/tempconv"
)

func main() {
	pointer() //from pointer.go
	echo()    //from echo.go

	c := tempconv.Celcius(220)
	k := tempconv.Kelvin(24)

	fmt.Println(c.String())
	fmt.Println(k.String())

	kel := tempconv.CtoK(c)
	fmt.Println(kel.String())

}
