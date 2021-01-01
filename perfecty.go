package main

import (
	"fmt"
	"github.com/rwngallego/perfecty-push/perfecty"
)

const filePath = "config/perfecty.yml"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Fatal error: %+v", e)
		}
	}()

	if err := perfecty.LoadConfig(filePath); err != nil {
		panic(err)
	}

	if err := perfecty.StartServer(); err != nil {
		panic(err)
	}
}
