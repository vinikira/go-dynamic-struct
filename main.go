package main

import (
	"fmt"
	api "go-dynamic-struct/API"
)

func main() {
	var (
		resp []api.Person
		err  error
	)

	if resp, err = api.GetPeople("http://localhost:8080/"); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
