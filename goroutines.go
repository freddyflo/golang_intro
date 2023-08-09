// A goroutine is a lightweight thread of execution.


package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")  // executes concurrently with the calling one -> f(s)

	go func(msg string) {
		fmt.Println(msg)
	} ("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}