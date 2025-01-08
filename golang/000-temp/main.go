package main

import "fmt"

func recoverFromPanic() {
    if r := recover(); r != nil {

		fmt.Println("Recovered. Error:\n", r)
	}
}

func pr() string {
	defer println("123")
	defer println("222")
	defer println("333")
	for{}
	return "abc"
}

func main() {
	go pr()
	select{}
}