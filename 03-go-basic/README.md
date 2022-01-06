# 03-go-basic
go 언어의 기본적인 함수를 알아봅시다!   
이번 과정에서는 함수 선언, 변수와 상수의 선언, return, 조건문과 반복문, 계산기를 만드는 작업을 해보겠습니다.

## 함수선언
![함수선언](./image/func.png)   
위와 같은 순서로 함수를 선언하면 됩니다.

## 변수선언

1. 변수를 하나 선언 (var 변수이름 타입)
```go
var num1 int
```
2. 같은 타입을 가지는 변수를 여러 개 선언
```go
var num2, num3 int
```
3. 여러 변수에 한 번에 값 초기화 : 선언과 동시에 값을 초기화하면 타입을 명시할 필요가 없음
```go
var num4, num5, str1 = 4, 5, "example"
```
4. 함수 안에서만 :=를 쓰면 var과 타입을 지정하지 않고 변수를 선언과 동시에 초기화 가능
```go
errorvar := str1
```
5. 다른 타입을 가지는 변수를 여러 개 선언
```go
var (
	i int
	b bool
	s string
)
```
위와같은 코드를 실행한다면,
```shell
$ go run var.go
```
output   
```shell
1:  0
2:  0 0
3:  4 5 example
4:  6
5:  0 false
```
string에 값을 지정하지 않으면 `""이 초기값`이 됩니다.


## 상수선언
변수 선언과 크게 다를 것이 없지만, 변수는 `var`를 이용했다면 상수는 `const`를 이용해서 선언합니다.   
1. 상수를 하나 선언 (const 변수이름 타입)
```go
const max_hp int = 100
const err_msg string = "잘못 사용했습니다"
```
2. 같은 타입을 가지는 상수를 여러 개 선언
```go
const max_x, max_y int = 100, 200
```
3. 여러 개의 상수를 블록 내에 정의해 선언
```go
const (
	start_x, start_y     int    = 0, 10
	msg_score, msg_level string = "스코어", "레벨"
	)
```
위와같은 코드를 실행한다면,
```shell
$ go run const.go
```
output   
```shell
100
잘못 사용했습니다
100
200
0
10
스코어
레벨
```


## print 
 fmt 패키지에서 제공하는 표준 출력 함수는 세가지가 존재합니다.  
 세가지의 특징을 살펴보겠습니다.  
   
- func Print(a ...interface{}) (n int, err error): 값을 그 자리에 출력(새 줄로 넘어가지 않음)
- func Println(a ...interface{}) (n int, err error): 값을 출력한 뒤 새 줄로 넘어감(개행)
- func Printf(format string, a ...interface{}) (n int, err error): 형식을 지정하여 값을 출력  

이 세가지의 출력함수를 살펴보기 위해 작성한 print파일을 실행해보면,
```shell
$ go run print.go
```
output
```shell
103.2Hello, world!
10
3.2
Hello, world!
정수: 10실수: 3.200000
문자열: Hello, world!
```

## return
1. return 뒤에 리턴 타입만 적어주는 방법
```go
func divide1(dividend, divisor int) (int, int) {
	var quotient = (int)(dividend / divisor)
	var remainder = dividend % divisor
	return quotient, remainder
}
```
2. return뒤에 리턴할 변수를 선언하는 방법   
1과는 달리 함수 내부에서 var로 선언하지 않아도 사용됩니다.   
마지막줄에 return이라고만 적으면 미리 return값으로 정해 놓은 quotient와 remainder를 return 됩니다.      
```go
func divide2(dividend, divisor int) (quotient, remainder int) {
	quotient = (int)(dividend / divisor)
	remainder = dividend % divisor
	return 
}
```
위와같은 코드를 실행한다면,
```shell
$ go run return.go
```
output   
```shell
1의 결과: 3 1
2의 결과: 3 1
```

## 조건문

#### IF문
if문의 기본 공식   
```
if 조건 {
    참일 때의 결과
}
```
#### IF-ELSE문
if문의 기본 공식   
참고로 `else`는 if문의 괄호 옆에 써야합니다.   
```
if 조건 {
    참일 때의 결과
} else {
    거짓일 때의 결과
}
```
위와같은 반복문을 담은 코드를 실행한다면,
```shell
$ go run if_else.go
```
output   
```shell
1.4142135623730951 음수는 x
9 20
45
```

#### SWITCH문
아래와 같은 형식을 이용하여 실행합니다.   
```go
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
```
  
- 다른 언어는 switch 키워드 뒤에 변수나 조건이 반드시 두지만, Go는 이를 `쓰지 않아도` 됩니다. 
- 이 경우 Go는 switch 조건을 true로 두고 첫번째 case문으로 이동하여 검사합니다.   
- 또한, `case문에 조건`을 둘 수 있습니다.   
- 다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는 특징이 있습니다.   

위와같은 코드를 실행한다면,
```shell
$ go build switch.go
$ ./switch.exe
```
output   
```shell
점수를 입력하세요
92
성적은 A 입니다
```

## 반복문
#### FOR문
1. 기본형태
for 키워드를 사용해 `초기화식, 조건식, 증감식`을 사용하여 반복문을 만듭니다.
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
fmt.Println(sum)
```
2. for문을 while문처럼 이용하기   
go언어에는 while문이 따로 존재하지 않기에 for문을 이용해서 만들어야 합니다.
```go
sum2 := 1
for sum2 < 100 {
    sum2 += sum2
}
fmt.Println(sum2)
``` 
3. 증감식을 for문 안에 넣어 사용하기
```go
sum3 := 0
j := 0
for j < 10 {
    sum3 += j
    j++
}
fmt.Println(sum3)
```
4. range를 이용하기
```go
alphabets := []rune{'A', 'B', 'C'}
for _, a := range alphabets {
    fmt.Println(a)
}
```
여기서 range는 각 원소의 `인덱스`와 `값`을 함께 반환합니다.   
위 예시에서는 인덱스값이 필요없기에 공백 식별자 _로 인덱스를 무시했습니다.   
또한 `rune`이라는 타입을 볼 수 있는데, 이는 **유니코드를 쉽게 제어하기 위한 타입**입니다. int32를 재정의하여 사용하기에 32bit로 제어가 됩니다.    
따라서, 직접 코딩을 할 때, "unicode/utf8"패키지에서 제공하는 API를 사용해 rune타입의 함수들을 쉽게 활용할 수 있다고 합니다.      

위와같은 코드를 실행한다면,
```shell
$ go run for.go
```
output   
```shell
45
128
45
65
66
67
```
## 계산기
위의 코드를 이용하여 간단한 계산을 할 수 있는 계산기를 만들어보았습니다.  
두개의 값을 받아 사칙연산을 하는 코드를 작성했습니다.
```shell
$ go build calculator.go
$ ./calculator.exe
```
output   
```shell
첫번쨰 숫자를 입력하세요
7
두번쨰 숫자를 입력하세요
4
입력하신 숫자는 7, 4 입니다 
연산자를 입력하세요+
7 + 4 = 11
```
