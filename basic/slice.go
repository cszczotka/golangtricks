package main

import (
	"fmt"
)

func main() {
	
	items := []string{"aaa", "bbb", "ccc"}
	for idx, item := range items {
		fmt.Println(idx, item)
	}
	
	for _, item := range items {
		fmt.Println(item)
	}
}
