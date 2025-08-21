package main

import (
	"main/handler"
	"main/service"
)

func main() {
	service.PrintFromService()
	handler.PrintFromHandler()
}
