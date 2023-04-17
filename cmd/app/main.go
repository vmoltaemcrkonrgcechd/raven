package main

import "raven/pkg/back"

func main() {
	b, err := back.ReadConfig("./.json")
	if err != nil {
		panic(err)
	}

	if err = b.Exec(); err != nil {
		panic(err)
	}
}
