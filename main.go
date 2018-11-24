package main

import (
	"fmt"
	"pkg/list"
)

func main() {
	a := []int64{1, 3, 5, 56, 67, 7}
	b := []int64{3, 5, 6, 7}
	fmt.Println(list.Sub(a, b))
}
