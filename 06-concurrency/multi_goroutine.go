package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(6) // 요고만 하면됨

	for i := 0; i < 100; i++ {
		go hello()
	}
}

func hello() {
	fmt.Println("hello")
}