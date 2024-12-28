package main

import (
	"fmt"
)

func testMe(x, y int, s string) {
	fmt.Println(x+y, s)
}

// public function in package start with upper case
func IntFuncWithError(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("Incorrect parameter value")
	}
	return 100, nil
}

func main() {
	fmt.Println("Running")
	testMe(1, 2, "aaa")
	myVar, err := IntFuncWithError(-1)
	if err != nil {
		panic(err)
	}
	fmt.Println(myVar)

}
