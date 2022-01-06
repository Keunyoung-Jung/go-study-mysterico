package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	sum3 := 0
	j := 0
	for j < 10 {
		sum3 += j
		j++
	}
	fmt.Println(sum3)

	alphabets := []rune{'A', 'B', 'C'}
	for _, a := range alphabets {
		fmt.Println(a)
	}

}
