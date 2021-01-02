package main

import (
	"fmt"
	"github.com/rwngallego/perfecty-push/perfecty"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Fatal error: %+v", e)
		}
	}()

	if err := perfecty.Load(); err != nil {
		panic(err)
	}

	if err := perfecty.StartServer(); err != nil {
		panic(err)
	}
}
