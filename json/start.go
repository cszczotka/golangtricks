package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"hero_name"`
	Email     string `json:"email"`
	Active    bool   `json:"is_active"`
}

func (a *User) activate() {
	a.Active = true
}

func main() {
	avengers := []User{
		{
			FirstName: "Bruce",
			LastName:  "Hulk",
			Email:     "bruce123@gmail.com",
		},
		{
			FirstName: "Tony",
			LastName:  "Stark",
			Email:     "midgard@dd.com",
		},
		{
			FirstName: "Thor",
			LastName:  "Odinson",
			Email:     "thoro@yahoou.com",
		},
	}

	avengers[1].activate()

	jsonBytes, err := json.Marshal(avengers)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
