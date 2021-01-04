package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, world!")
	// 4
	fmt.Println(Inc())
}

func Inc() (v int) {
	defer func() {
		v++
	}()
	return 3
}
