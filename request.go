package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("err=", err)
		panic("We got into some serious problem")
	}
	res := &[]Response{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&res); err != nil {
		fmt.Println("err=", err)
		panic("We got into some serious problem")
	}
	for _, value := range *res {
		fmt.Println("id=", value.Id)
		fmt.Println("userid=", value.UserId)
		fmt.Println("title=", value.Title)
	}
}
