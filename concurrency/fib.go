package main

import (
	"fmt"
)

func main() {
	var limit int = 50
	jobs := make(chan int, limit)
	results := make(chan int, limit)

	go worker(jobs, results)

	for i := 0; i<limit ; i++ {
		jobs <- i
	}

	close(jobs)

	for j :=0; j < limit; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}