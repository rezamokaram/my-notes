package main

type Sample interface {
}

type MyStruct struct {
}

func main() {
	var inf Sample
	if inf == nil {
		println("true")
	} else {
		println("false")
	}

	inf = MyStruct{}
	if inf == nil {
		println("true")
	} else {
		println("false")
	}
}