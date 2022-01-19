package main

import "fmt"

func routine1(channel chan int){
	x := <- channel
	fmt.Println(x)
	if x == 0 {
		channel <- x+10
	} else {
		channel <- x+60
	}
}

func routine2(channel chan int){
	x := <- channel
	fmt.Println(x)
	if x == 10 {
		channel <- x*10
	} else {
		channel <- x-10
	}
}

func main() {
  // 채널 생성
	myChannel := make(chan int)
	go func(){
		x := 0
		myChannel <- x
		} ()
	go routine1(myChannel)
	go routine2(myChannel)
	go routine1(myChannel)
	go routine2(myChannel)
	fmt.Scanln()
	x := <- myChannel
	fmt.Println(x)
}