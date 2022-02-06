package main

import (
	"errors"
	"fmt"
)

func plus(n1, n2 int) (result int) {
	result = n1 + n2
	return
}
func minus(n1, n2 int) (result int) {
	result = n1 - n2
	return
}
func multi(n1, n2 int) (result int) {
	result = n1 * n2
	return
}
func divide(n1, n2 int) (result int, err error) {
	if n1 < 0 || n2 < 0 {
		return 0, errors.New("음수는 나누기가 불가능합니다")
	} else if n1 == 0 || n2 == 0 {
		return 0, fmt.Errorf("0은 나누기가 불가능합니다")
	} else {
		result = n1 / n2
		return result, nil
	}
}

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
		fmt.Println(plus(n1, n2))
	} else if line == "-" {
		fmt.Println(minus(n1, n2))
	} else if line == "*" {
		fmt.Println(multi(n1, n2))
	} else if line == "/" {
		if f, err := divide(n1, n2); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println(f)
		}
	} else {
		fmt.Printf("잘못 입력하셨습니다")
	}

}
