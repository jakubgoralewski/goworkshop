package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Jacob", Age: 24}

	bytes, _ := json.Marshal(u)
	fmt.Println(string(bytes))
}
