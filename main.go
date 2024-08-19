package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6}

	a := data[1:3:4]

	fmt.Println(a, len(a), cap(a))
}
