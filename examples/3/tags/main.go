package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
}

func main() {
	b := []byte(`{"name": "username", "password": "pwd"}`)
	var u User
	if err := json.Unmarshal(b, &u); err != nil {
		panic(err)
	}

	fmt.Println(u)
	b2, _ := json.Marshal(u)
	fmt.Println(string(b2))
}
