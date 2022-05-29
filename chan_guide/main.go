package main

import (
	"fmt"
)

func main() {
	ch := make(chan interface{}, 1)
	//ExecuteTimeout := time.Second * 3
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("==============")
	//		//ch <- struct{}{}
	//	}
	//}()
	//竟然会阻塞住！！！
	select {
	case r := <-ch:
		fmt.Println("chchchchchchchchchchchchch", r)
	//case <-time.After(ExecuteTimeout):
	//	fmt.Println("ExecuteTimeout")
	}

	fmt.Println("main done")

}
