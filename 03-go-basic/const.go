package main

import "fmt"

func main() {
	const max_hp int = 100
	const err_msg string = "잘못 사용했습니다"

	fmt.Println(max_hp)
	fmt.Println(err_msg)

	const max_x, max_y int = 100, 200 //여러 개의 상수 정의
	fmt.Println(max_x)
	fmt.Println(max_y)

	const (
		start_x, start_y     int    = 0, 10
		msg_score, msg_level string = "스코어", "레벨"
	) //상수 여러 개를 괄호 블럭 내에 정의
	fmt.Println(start_x)
	fmt.Println(start_y)
	fmt.Println(msg_score)
	fmt.Println(msg_level)
}
