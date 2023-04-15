package main

import (
	"fmt"
	"raven/pkg/back"
)

func main() {
	if err := back.ReadConfig("./join.json"); err != nil {
		fmt.Println(err)
	}
}
