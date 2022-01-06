package main

import "fmt"

func return_grade(score int) {
	var grade string
	switch {
	case score > 90:
		grade = "A"
	case score > 70:
		grade = "B"
	case score > 50:
		grade = "C"
	default:
		grade = "F"
	}

	fmt.Println("성적은", grade, "입니다")
}

func main() {
	fmt.Println("점수를 입력하세요")
	var score int
	fmt.Scanln(&score)

	return_grade(score)
}
