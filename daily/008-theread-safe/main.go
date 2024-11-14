package main

import(
	// "sync"
)

func main() {
    // a := sync.Map{
	// 	mu: new(sync.Mutex)
	// }

	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		a[1]++
	// 	}()
	// }

    select{} // block-forever trick
}

// fatal error: concurrent map iteration and map write