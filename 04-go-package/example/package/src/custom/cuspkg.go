package cuspkg

import (
	"fmt"
	"runtime"
)

func Show() {
	fmt.Println(runtime.GOOS)
}