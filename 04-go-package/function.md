# 04 - Go Package

Go 언어의 문법과 함수 등에 대해 좀 더 알아봅시다 🏌️‍♀️


<br /><br />

## Go의 내장 함수

Go의 [내장함수](https://thebook.io/006806/ch02/04/04/)에 대해서 알아보자. 

이 내용은 문법에 대하여 세세한 명세는 달지 않는다. 문법상으로 특이한 케이스가 아니면 간략하게 설명하며, 혹은 생략하기도 한다.

만일 코드 작성 중 문법에 관하여 어려운 점이나 더 배워야 할 필요성을 느끼는 경우 [예제로 배우는 Go 프로그래밍](http://golang.site/go/basics)이나 [Go 언어 웹 프로그래밍 철저 입문](https://thebook.io/006806/) 등의 책을 참고하길 바란다.

<br />

### Go 콜렉션 함수

Go에서 일컫는 컬렉션은 `배열`, `슬라이스`, `맵`으로 나뉜다.

- `배열`은 연속된 메모리 공간에 동일한 타입의 데이터를 순차적으로 저장한다.
	1. Go에서 배열의 크기는 타입으로서의 의의를 가진다. 때문에 `[3]int`와 `[5]int`는 동일하지 않은 타입으로 인식한다.
	2. 다차원 배열을 지원한다.
	```
	var multiArray [3][4][5]int  // 정의
	multiArray[0][1][2] = 10     // 사용
	```

- `슬라이스`는 기본적으로 배열과 유사한 구조이나, 개발 시 편리한 기능을 몇 가지 가지고 있다.
	1. 크기를 동적으로 변경할 수 있다.
	2. 배열의 일부를 발췌하거나 변경하는 것이 가능하다.

- `맵`은 `Key`와 `Value`를 가지고 있는 형태이다. [해시테이블](https://mangkyu.tistory.com/m/102)을 구현한 형태이기 때문에 키에 대응하는 값을 빠르게 찾을 수 있다.



#### `append()`

슬라이스에 값을 순차적으로 저장한다.

```
package main
 
import "fmt"
 
func main() {
    // len=0, cap=3 인 슬라이스
    sliceA := make([]int, 0, 3)
 
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(sliceA), cap(sliceA))
    }
 
    fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력 
}
```

<br />

#### `copy()`

슬라이스에 저장된 요소를 복사할 수 있다.

```
func main() {
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    fmt.Println(target)  // [0 1 2 ] 출력
    println(len(target), cap(target)) // 3, 6 출력
}
```

<br />

#### `len()`

슬라이스의 원소 개수(즉, 배열의 길이)를 확인할 수 있다.

<br />

#### `cap()`

슬라이스의 전체 크기(최대 용량)를 확인할 수 있다.

```
func main() {
    s := make([]int, 5, 10)
    println(len(s), cap(s)) // len 5, cap 10
}
```


<br />

---

### Go 복소수 함수

#### `complex`
실수부/허수부를 나타낼 수 있는 타입니다.
뒤에 붙은 숫자를 통해 비트수를 알 수 있다. 

ex) `complex64`(32비트), `complex128`(64비트)

#### `real()`

complex 타입 변수의 실수값에 해당하는 부분을 가져온다.

#### `imag`

complex 타입 변수의 허수값에 해당하는 부분을 가져온다.

```
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```

<br />

---

<br /><br />

## 클로저(Closure)


다음 [예제](https://go.dev/tour/moretypes/25)를 확인해보자.

```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
		)
	}
}
```

### **클로저**란 뭐지?

클로저에 대해서는 다양한 정의가 있으나 `어떤 함수에서 선언한 변수를 참조하는 내부함수 B를 외부로 전달할 경우, 실행 컨텍스트가 끝난 후에도 변수가 사라지지 않는 현상` 이라고 요약하겠다.

동작 방식에 따라 차이가 있지만 일반적으로 실행 컨텍스트가 끝난 언어는 더이상 사용되지 않는 것으로 확인된다. 그리고 곧 가비지 컬렉터에 따라 메모리가 해제된다.

그러나 클로저의 경우에는 다르다. 함수를 return 값으로 전달할 때 그 함수가 어떤 값을 참조하고 있는 경우 가비지 컬렉터는 대상을 수집하지 않는다. 그러므로 실행 컨텍스트가 끝난 후에도 계속 유지된다.

이때, `Closure`를 지원하는 언어는 하나의 조건을 요한다. 함수가 **일급 객체**로 다뤄져야 한다는 것이다.

<br />

### **일급 객체**가 뭐지?

- 정수
- 실수
- 일부 문자열

일급 객체는 함수의 인자가 되고, 반환값이 되고, 수정이나 할당이 가능하다.  
또한 변수에 대입하는 것 또한 가능하다.

파이썬은 함수를 *일급 객체*로 다룬다. 때문에 Closure 기능을 지원한다.

```
def add(a, b):
    return a + b

def execute(func, *args):
    return func(*args) #*args는 여러 인수를 받는다는 의미

f = add

>>> execute(f, 3, 5) # 1.

8
```

다음 예제에서는 execute라는 함수로 add 함수를 한 번 감쌌다.

이 과정에서 f라는 함수는 execute의 인자로 전달되었다.

또한 add라는 함수는 f에 저장되었다.

<br />

자바스크립트 또한 함수를 *일급 객체*로 다룬다.

```

const func_len = (data) => {
    return data.length;
}

const spliter = (arr) => {
    return arr[0];
}

let data = ["hello", "world", "this", "is", "study"];

const value = func_len(spliter(data));
//value : 5

```

함수를 변수에 대입하고, 인자에 보내고, 그 값을 받아오는 일련의 과정이 자연스럽게 이루어진다. 함수를 `일급 객체`로 다룬다는 의미이다.


Go에서의 Closure를 보고 넘기도록 하자.
```
package main
  
import (
    "fmt"
    "strings"
)
  
func makeSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}
  
func main() {
    addZip := makeSuffix(".zip")
    addTgz := makeSuffix(".tar.gz")
    fmt.Println(addTgz("go1.5.1.src"))
    fmt.Println(addZip("go1.5.1.windows-amd64"))
}
```

<br /><br />

## 함수를 매개변수로 전달하는 법

Go가 함수를 일급 객체로 다룬다는 사실이 밝혀졌으니, 함수를 매개변수로 전달하는 일이 이제 와서 그다지 새삼스럽지 않을 것이다.

간단한 예제를 보고 넘기도록 하자.

```
package main
  
import "fmt"
  
func callback(y int, f func(int, int)) {
    f(y, 2) // add(1, 2)를 호출
}
  
func add(a, b int) {
    fmt.Printf("%d + %d = %d\n", a, b, a+b) // 1 + 2 = 3
}
  
func main() {
    callback(1, add) // 1 + 2 = 3
}

```

<br /><br />

## 참고 자료

http://golang.site/go/article/11-Go-%ED%81%B4%EB%A1%9C%EC%A0%80

https://cjwoov.tistory.com/8

https://thebook.io/006806/ch02/04/03/

https://thebook.io/006806/ch03/02/03/

https://go.dev/tour/moretypes/25

https://hwan-shell.tistory.com/339

https://ko.wikipedia.org/wiki/%ED%81%B4%EB%A1%9C%EC%A0%80_(%EC%BB%B4%ED%93%A8%ED%84%B0_%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D)

https://shoark7.github.io/programming/python/closure-in-python

https://ko.wikipedia.org/wiki/%EC%9D%BC%EA%B8%89_%EA%B0%9D%EC%B2%B4

https://velog.io/@graphicnovel/JS-%ED%81%B4%EB%A1%9C%EC%A0%80%EC%9D%98-%EC%9D%98%EB%AF%B8%EC%99%80-%EC%9B%90%EB%A6%AC-%EC%9D%B4%ED%95%B4-Closure
