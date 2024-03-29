package main

import (
	"fmt"
	"github.com/perfectyorg/perfecty-push"
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
