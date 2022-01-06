package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 float32 = 3.2
	var s string = "Hello, world!"

	// Print: 값을 그 자리에 그대로 출력, 줄넘기지않음
	fmt.Print(num1)    // 10: 정수 출력
	fmt.Print(num2)    // 3.2: 실수 출력
	fmt.Print(s, "\n") //  Hello, world!: 문자열 출력

	// Println: 값을 출력한 뒤 새 줄로 넘어감
	fmt.Println(num1) // 10: 정수 출력
	fmt.Println(num2) // 3.2: 실수 출력
	fmt.Println(s)    //  Hello, world!: 문자열 출력

	fmt.Printf("정수: %d", num1)   // 정수: 10
	fmt.Printf("실수: %f\n", num2) // 실수: 3.2
	fmt.Printf("문자열: %s\n", s)   // 문자열: Hello, world!

}
