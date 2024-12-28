package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("aaa", c)

	/*
		for {
			msg, open := <- c
			if !open {
				break
			}
			fmt.Println(msg)
		}
	*/
	for msg  := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 1000)
	}
	close(c)
}
