package main

import "testing"

var (
	data []string
)

func TestFirst(t *testing.T) {
	Operation()
}

func TestSecond(t *testing.T) {
	Operation()
}

func Operation() {
	data := append(data, "string_1")
	for _, item := range data {
		print(item, " ")
	}
	println()
}

// func main() {
// 	println("hi")
// }