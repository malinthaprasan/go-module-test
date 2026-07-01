package main

import (
	"fmt"
	"hello-again/svc"
)

func main() {
	fmt.Println(svc.Hello())
}

func HelloFromMain() string {
	return "Hello Main"
}
