package main

import (
	"fmt"
	"hello-again/svc"

	"github.com/malinthaprasan/go-module-test/another-module-in-base/svc"
)

func main() {
	fmt.Println(svc.Hello())
}

func HelloFromMain() string {
	return "Hello Main"
}
