package main

import (
	"fmt"
)

func main() {
	x := 50
	y := &x
	fmt.Println(&x,":",x)
	fmt.Println(&y,":",y)
	fmt.Println(&*y,":",*y)
	*y = 100
	fmt.Println("----*y = 100 선언----")
	fmt.Println(&x,":",x)
}
