package main

import "fmt"

// 1 변수를 하나 선언 (var 변수이름 타입)
var num1 int

// 2 같은 타입을 가지는 변수를 여러 개 선언
var num2, num3 int

// 3 여러 변수에 한 번에 값 초기화 : 선언과 동시에 값을 초기화하면 타입을 명시할 필요가 없음!
var num4, num5, str1 = 4, 5, "example"

// 4 함수 안에서는 :=를 쓰면 var과 타입을 지정하지 않고 변수를 선언과 동시에 초기화 가능
//하지만 함수 밖에서는 :=를 쓸 수 없음 따라서 아래 줄은 에러
//errorvar := str1

// 5 다른 타입을 가지는 변수를 여러 개 선언
var (
	i int
	b bool
	s string
)

//출력
func main() {
	fmt.Println("1: ", num1)
	fmt.Println("2: ", num2, num3)
	fmt.Println("3: ", num4, num5, str1)

	//4 함수 안에서 선언
	num6 := 6
	fmt.Println("4: ", num6)
	fmt.Println("5: ", i, b, s)
}
