package main

import "fmt"

func main() {
	a := "\x48\x65\x6C\x6C\x6F\x3A\x2D\x29"

	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%X ", a[i])
		fmt.Printf("%b\n", a[i])

	}
}
