package main

import "fmt"

//1 return 뒤에 리턴 타입을 적어주는 방법
func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}

//2 return뒤에 리턴할 변수를 선언하는 방법. 1과는 달리 함수 내부에서 quotient를 var로 선언하지 않고 바로 사용
func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return //return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return 됨
}

func main() {
	//1
	var quotient, remainder int
	quotient, remainder = divide1(10, 3)
	fmt.Println("1의 결과:", quotient, remainder)

	//2
	var quotient2, remainder2 = divide2(10, 3)
	fmt.Println("2의 결과:", quotient2, remainder2)
}
