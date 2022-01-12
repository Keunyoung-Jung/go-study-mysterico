# 04 - Go Package

Go 언어의 문법과 함수 등에 대해 좀 더 알아봅시다 🏌️‍♀️
<br /><br />

## Go의 내장 함수

### Go 콜렉션 함수

#### `append()`

슬라이스에 값을 순차 보관한다.

#### `copy()`

슬라이스에 저장된 요소를 복사할 수 있다.

#### `delete()`

슬라이스의 요소를 삭제할 수 있다.

#### `len()`

슬라이스의 원소 개수(즉, 배열의 길이)를 확인할 수 있다.

#### `cap()`

슬라이스의 전체 크기를 확인할 수 있다.

```
여기에 예시
(append, copy, delete, len, cap이 모두 들어가야함!)
```

[링크](example/function/src/collection.go)

<br />

---

### Go 복소수 함수

#### `complex()`

#### `real()`

#### `imag()`

```
여기에 예시
(complex, real, imag가 모두 들어가야함!)
```

[링크](example/function/src/complicated.go)

<br />

---

### Go 오류 관리 함수

#### `defer`

#### `panic`

#### `recover`

```
여기에 예시
(defer, panic, recover)
```

[링크](example/function/src/error.go)

<br /><br />

## 클로저(Closure)

클로저는 바깥의 변수를 함수 안으로 끌어들이는 것처럼 참조하는 함수의 값을 이르는 말이다.

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

(대충 여기에 클로저가 뭔지 이야기하기)

이때, `Closure`를 지원하는 언어는 하나의 조건을 요한다. 함수가 **일급 객체**로 다뤄져야 한다는 것이다.

<br />

### **일급 객체**가 뭐지?

- 정수
- 실수
- 일부 문자열(지원하지 않는 경우 있음)

일급 객체는 함수의 인자가 되고, 반환값이 되고, 수정이나 할당이 가능하다.  
또한 변수에 대입하는 것 또한 가능하다.

파이썬은 함수를 일급 객체로 다룬다. 때문에 Closure 기능을 지원한다.

```
def add(a, b):
    return a + b

def execute(func, *args):
    return func(*args) # 2.

f = add # 3.

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

<br /><br />

## 함수를 매개변수로 전달하는 법

Go가 함수를 일급 객체로 다룬다는 사실이 밝혀졌으니, 함수를 매개변수로 전달하는 일이 이제 와서 그다지 새삼스럽지 않을 것이다.

간단한 예제를 보고 넘기도록 하자.

```

```

[예제]()

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
