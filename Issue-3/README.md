# [Issue] 외부함수 호출하는 방법 #3

해당 내용은 **외부함수 호출하는 방법** 이슈에 대한 해결을 위한 branch입니다.   
   
이슈는 [여기 링크](https://github.com/Keunyoung-Jung/go-study-mysterico/issues/3) 에서 확인할 수 있습니다.   

## Questions
```Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("점수를 입력하세요")
	var score int
	fmt.Scanln(&score)

	Return_grade(score)
}
```
이것이 main.go의 내용이고
아래의 것이 Return_grade함수가 있는 코드는 아래의 것입니다.
```Go
package Grade_switch

import "fmt"

func Return_grade(score int) {
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
```
실행 시켜주실분 누구일까요??ㅎㅎㅎㅎ

## Answer
