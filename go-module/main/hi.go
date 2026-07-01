package main

import (
	"fmt"
	"hello-again/svc"

	exp "github.com/malinthaprasan/go-module-test/another-module-in-base/svc"
	expfol "github.com/malinthaprasan/go-module-test/another-module-in-folder/this-is-module/folder-svc"
)

func main() {
	fmt.Println(svc.Hello())
	fmt.Println(exp.Hello())
	fmt.Println(expfol.Hello())
}

func HelloFromMain() string {
	return "Hello Main"
}
