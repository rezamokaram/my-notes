package service

import "main/handler"

func PrintFromService() {
	println("Hello from service")
	handler.PrintFromHandler()
}
