package main

import "fmt"

func main() {

	a := []int{0, 1, 2, 3, 4}

	aa := a
	a[2] = -888
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &aa)
	//修改a 会影响aa
	fmt.Println(a)
	fmt.Println(aa)

}
