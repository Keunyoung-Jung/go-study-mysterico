package main

import (
	"fmt"
)

func main() {

	fmt.Println("첫번쨰 숫자를 입력하세요")
	var n1 int
	fmt.Scanln(&n1)

	fmt.Println("두번쨰 숫자를 입력하세요")
	var n2 int
	fmt.Scanln(&n2)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다 \n", n1, n2)

	fmt.Printf("연산자를 입력하세요")
	var line string
	fmt.Scanln(&line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%d x %d = %d", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Printf("잘못 입력하셨습니다")
	}

}
