package main

import (
	"cli-app/isprime"
	"fmt"
)

func main() {
	n := 2

	_, msg := isprime.IsPrime(n)

	fmt.Println(msg)
}
