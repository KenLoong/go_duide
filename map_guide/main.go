package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,3,4,6,8,9}
	a = append(a[:len(a)-1])
	fmt.Println(a)

}

func f1() {
	m1 := make(map[int]string)
	m2 := make(map[int]string)

	for i := 0; i < 6; i++ {
		m1[i] = fmt.Sprintf("%v", i)
	}

	for i := 4; i < 8; i++ {
		m2[i] = fmt.Sprintf("%v", i)
	}

	filterExistedKey(m1, m2)
	fmt.Println(m2) //打印结果：map[6:6 7:7]
}

//是会直接影响m2的
func filterExistedKey(m1 map[int]string, m2 map[int]string) {
	for key, _ := range m2 {
		if _, ok := m1[key]; ok {
			delete(m2, key)
		}
	}
}
