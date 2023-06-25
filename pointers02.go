package main

import "fmt"

func main() {
	a := 3
	var ptr *int = &a
	parOuImpar(ptr)
	fmt.Println(a)

}
func parOuImpar(ptr *int) int {

	if *ptr%2 != 0 {
		*ptr = 1
		return *ptr
	}
	*ptr = 0
	return *ptr
}
