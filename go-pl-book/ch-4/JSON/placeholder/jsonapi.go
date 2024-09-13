package placeholder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "https://jsonplaceholder.typicode.com/users"

type User struct {
	id       int
	name     string
	username string
	email    string
	// address  Address
}

type Address struct {
	street  string
	city    string
	zipcode string
}

var userList []User

func GetUserData() {
	resp, err := http.Get(endpoint)

	if err != nil {
		resp.Body.Close()
		fmt.Println("query failed")
	}

	// NewDecoder() creates a new Decoder which reads data from whatever value was passed
	err = json.NewDecoder(resp.Body).Decode(&userList)
	if err != nil {
		resp.Body.Close()
		fmt.Println("Error decoding JSON data")
	}

	fmt.Println(userList)
	defer resp.Body.Close()
}
