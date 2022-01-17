# 05 Data types
## Boolean
* true
* false
```Go
b1 := true
```
## Integer
* int, uint (int,uint 만 쓰는 경우 머신의 비트에 따름)
* int8, uint8
* int16, uint16
* int32, uint32
* int64, uint64
* byte (== uint8)
* rune (== int32)
* uintptr (== uint / 포인터를 저장할때 사용)
```Go
a := 365    // 10진수
b := 0555   // 8진수
c := 0x160  // 16진수
```
## Float
* float32
* float64
```Go
a := 4.12
b := 1.2345e-3
```
## Complex
* complex64
* complex128
```Go
a := 1+2i
b := complex(5,6)
```
## Math function
수학함수는 너무너무 많아서 자주 쓸것 같은것만 작성
* math.Cos(x)
* math.Dim(x)
* math.