package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("a = %v and the ref = %v\n", a, &a)

	a = 2
	fmt.Printf("a = %v and the ref = %v\n", a, &a)

	// // mutable
	// var x []int = []int{1, 2, 3, 4, 5}
	// y := x
	// y[0] = 100
	// fmt.Printf("x = %v\n", x)
	// fmt.Printf("y = %v\n", y)

	//  // immutable
	//  var X = 5
	//  Y := X
	//  Y = 7
	//  fmt.Printf("x = %v\n", X)
	//  fmt.Printf("y = %v\n", Y)

}
