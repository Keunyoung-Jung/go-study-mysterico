package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("C:\\workspace\\go-study-mysterico\\README.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	// println(f.Name())
}
