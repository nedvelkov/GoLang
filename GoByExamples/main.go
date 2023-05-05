package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	user := unmarshalJsonString("{\"firstName\":\"Bob\",\"lastName\":\"Evans\"}")
	if user == nil {
		fmt.Println("User is nil")
	} else {

		fmt.Println(user.FirstName)
	}
}

func unmarshalJsonString(text string) *User {
	var user User

	if err := json.Unmarshal([]byte(text), &user); err != nil {
		fmt.Println(err)
		return nil
	}
	return &user
}
