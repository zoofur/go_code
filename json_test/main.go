package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("json_test/json.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	u := new(user)
	err = json.Unmarshal(data, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
	fmt.Println(u.Data[123123])
}

type user struct{
	Name string
	Data map[int]interface{}
}
