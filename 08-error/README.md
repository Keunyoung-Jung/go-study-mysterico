# 08 Error
이번 장에서는 에러처리에 대해서 설명합니다.

## 에러와 예외
둘은 비슷하지만 완전히 같은 의미는 아닙니다.

우선, 공통점은 발생하면 프로그램이 정상진행이 불가피하여 프로그램을 `종료`한다는 점입니다.   
그리고 과거에는 예외처리가 따로 존재하지 않았을 때는 모두 에러라고도 불렀다고 합니다.

하지만 가장 큰 `차이점`은 회피가 가능한지, 불가능하지로 분류됩니다.

`오류`의 경우, 시스템이 종료되어야 할 수준의 상황과 같이 수습할 수 없는 심각한 문제를 의미해 개발자가 미리 예측 방지가 `불가능`, 따라서 `회피가 불가능`합니다.   
반면 `예외`의 경우, 개발자가 구현한 로직에서 발생한 실수나 `사용자의 영향`에 의해 발생하기 떄문에 오류와 달리 개발자가 미리 예측하여 방지가 가능 즉, `회피가 가능`합니다.

## 에러 타입
Go에서는 기본으로 제공하는 error 라는 interface 타입을 가집니다.   
아래와 같이 이용하면 됩니다. 즉, Error() string 이라는 메서드를 가집니다.
```go
type error interface
{
    Error() string
}
```

만약 함수가 결과와 에러를 함께 리턴한다면, 이 에러가 nil 인지를 체크해서 에러가 존재여부를 체크할 수 있습니다.    
예를 들어, error를 체크해서 값이 `nil 이면 에러가 없는 것`이고, nil 이 아니면 err.Error() 로부터 해당 에러를 확인할 수 있습니다.    
아래 예제는 값이 1이 들어올 때는 정상동작을 하고, 그 외의 값이 들어올 때는 에러를 발생시키는 예제입니다. (log패키지의 Fatal 함수는 에러 로그를 출력하고 프로그램을 완전히 종료)   
```go
func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}
}
```
[작동시키기](https://go-tour-ko.appspot.com/welcome/1)

Go의 내장 함수나 기본 라이브러리의 대부분의 함수와 메서드는 에러가 발생할 수 있다면 `error 값을 함께 반환`하도록 작성되어 있습니다.    


## 에러 생성
의도적으로 에러를 발생시키려 할 때 가장 간단히 할 수 있는 방법은 `errors.New()` 를 사용해서 에러를 생성하는 방법입니다.    
또한  `fmt.Errorf()`도 동일하게 사용이 가능합니다.   
이는 cal-error의 예제를 통해 구현했습니다.

```go
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
```

## 사용자 정의 에러타입
만약 앞서 소개한 에러보다 더 다양한 기능을 반환하는 에러를 만들거나, 자신이 만든 메서드에서만 반환하는 특별한 에러가 존재하는 경우 직접 정의한 에러를 반환시킬 수 있습니다.   

1. 에러를 정의
아래는 SqrtError 예시로 에러 발생시, 에러가 발생한 시간, 발생시킨 값, 에러메세지를 한번에 출력시키는 에러를 만드는 방법입니다.
```go
type SqrtError struct {
	time    time.Time // 에러가 발생한 시간
	value   float64   // 에러를 발생시킨 값
	message string    // 에러 메시지
}

// error 인터페이스에 정의된 Error() 메서드 구현
func (e SqrtError) Error() string {
	return fmt.Sprintf("[%v]ERROR - %s(value: %g)", e.time, e.message, e.value)
}

func Sqrt(f float64) (float64, error) {
    // 매개변수로 전달된 값이 유효한 값이 아닐 때 SqrtError를 반환
    if f < 0 {
        return 0, SqrtError{time: time.Now(), value: f, message: "음수는 사용할 수 없습니다"}
    }
    if math.IsInf(f, 1) {
        return 0, SqrtError{time: time.Now(), value: f, message: "무한대 값은 사용할 수 없습니다"}
    }
    if math.IsNaN(f) {
        return 0, SqrtError{time: time.Now(), value: f, message: "잘못된 수입니다"}
    }
     
    // 정상 처리 결과 반환
    return math.Sqrt(f), nil
}
```

output
```shell
2022/02/06 17:02:17 [2022-02-06 17:02:17.6286128 +0900 KST m=+0.015342901]ERROR - 음수는 사용할 수 없습니다(value: -9)
exit status 1
```

2. 에러 타입별로 별도의 에러를 처리하는 방식
코드를 작성할 때 에러 처리를 더 명확히 알 수 있도록 에러르 변수화 하여 처리합니다.
```go
var (
    ErrA = errors.New("Error A")
    ErrB = errors.New("Error B")
    ...
    ErrZ = errors.New("Error B")
)
```
위에 선언한 변수들은 에러에 대한 컨텍스트를 담게 되고 외부에서 해당 변수들을 타입 검증을 통해 에러 처리를 할 수 있습니다.   
주로 switch문을 이용해서 처리합니다.
```go
func main() {
	if err := something(); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad request occured")
			return
		case ErrPageMoved:
			fmt.Println("The Page moved")
			return
		default:
			fmt.Println(err)
			return
		}
	}
}
```

## 패닉(panic)
go에는 에러와 패닉이 존재하는데, 패닉은 런타임 중에 발생하는치명적인 에러를 의미합니다. 죽, 프로그램이 더 이상 진행될 수 없는 수준의 에러를 뜻하고 발생하면 프로그램이 종료됩니다. (일반적인 에러로는 프로그램이 종료되지 않습니다.)   
패닉의 예로는 배열의 범위를 넘어서는 인덱스로 접근 또는 타입 어설션 실패가 있다고 합니다. 
```go
package main

import "fmt"

func main() {
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}
```
output
```shell
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.divide(...)
        C:/workspace/go-study/src/08-error/panic.go:10
main.main()
        C:/workspace/go-study/src/08-error/panic.go:6 +0x12
exit status 2
```