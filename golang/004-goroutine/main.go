package main

import "time"

func main() {
	data := make([]string, 0)
	data = append(data, "string_1")
	go printAfter(data, time.Second*1)

	data = append(data, "string_2")
	go printAfter(data, time.Second*1)

	println("----------------------------------------")
	println("my length is: ", len(data))
	println("my addr is: ", &data)
}

func printAfter(data []string, d time.Duration) {
	time.Sleep(d)
	println("----------------------------------------")
	println("my length is: ", len(data))
	println("my addr is: ", &data)
}
