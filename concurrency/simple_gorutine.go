package main
import (
	"fmt"
	"time"
 
)

func main() {
	//count("aaa")
	//count("bbb")
	go count("aaa")
	go count("bbb")

	fmt.Scanln()
}

func count(thing string) {
	for i:=1; true; i ++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 1000)
	}
}