package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

func main () {
	in := `{"FirstName":"Joshn","lastName":"Dow"}`
	bytes := []byte(in)

	var p Personerr := json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	fmt.Print("%+v",p)
}
