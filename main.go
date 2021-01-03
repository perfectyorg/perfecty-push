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

	if err := perfecty.Start(); err != nil {
		panic(err)
	}
}
