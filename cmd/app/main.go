package main

import (
	"fmt"
	"raven/pkg/back"
)

func main() {
	if err := back.ReadConfig("./.json"); err != nil {
		fmt.Println(err)
	}
}
