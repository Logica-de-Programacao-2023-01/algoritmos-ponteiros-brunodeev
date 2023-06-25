package main

import (
	"fmt"
	"math"
)

func mediaPi(n *float64) {
	*n = (math.Pi + *n) / 2
}

func main() {
	x := 25.92
	mediaPi(&x)
	fmt.Println(x)
}
