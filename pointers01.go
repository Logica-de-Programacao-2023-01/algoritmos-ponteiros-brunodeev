package main

import "fmt"

func main() {
	x := 1
	n := 4
	var ptr *int = &x
	var ptr2 *int = &n
	fmt.Print(Sum(ptr, ptr2))
}

func Sum(ptr *int, ptr2 *int) int {
	soma := 0
	for i := *ptr; i <= *ptr2; i++ {
		soma += i
	}

	return soma
}
