package main

import "fmt"

func main() {
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	prev := 3
	a = a[:prev]
	fmt.Println(a)
}
