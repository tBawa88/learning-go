package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID       int
	Name     string
	Addr     string
	DOB      time.Time
	Position string
	Salary   int
}

func main() {
	var eList []Employee
	eList = append(eList, Employee{ID: 1, Name: "Tarun", Addr: "India", DOB: time.Now(), Position: "Senior Dev", Salary: 10000000})

	for _, emp := range eList {
		fmt.Printf("%+v\n", emp)
	}

}
