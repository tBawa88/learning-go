package main

import "fmt"

func pointer() {
	x := 12
	p := &x
	fmt.Println(*p)

	incr(&x)
	fmt.Println(incr(&x)) //14

	ptr := new(int)
	*ptr++
	fmt.Printf("%d is value of new pointer", *ptr)
	incr(p)
}

func incr(ptr *int) int {
	*ptr++
	return *ptr
}
