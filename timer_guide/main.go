package main

import (
	"fmt"
	"os"
	"time"
)

const timeCount = time.Millisecond * 100

var myTimer *time.Timer

func main() {

	//testReadAndStop()
	//testResetAndStop()

	c := make(chan struct{})
	<-c
}

func testReadAndStop() {
	//先开始定时器
	myTimer = time.NewTimer(timeCount)
	go readTimer()
	go stopTimer()
}

func testResetAndStop() {
	//先开始定时器
	myTimer = time.NewTimer(10 * timeCount)
	go resetTimer(10 * timeCount)
	go func() {
		//给时间去reset timer
		time.Sleep(2 * timeCount)
		for true {
			select {
			case <-myTimer.C:
				//永远也不会触发，适用于心跳场景，当fire时，就自己开始选举
				fmt.Printf("fire !!!\n")
				os.Exit(1)
			}
		}
	}()
}

func readTimer() {
	//让stopTimer先执行
	time.Sleep(3 * timeCount)
	<-myTimer.C
	fmt.Printf("readTimer read from my Timer\n")
}

func resetTimer(count time.Duration) {
	for true {
		time.Sleep(timeCount)
		//永远也不会触发，适用于心跳场景，当fire时，就自己开始选举
		ok := myTimer.Reset(count)
		fmt.Printf("resetTimer Reset Timer get %v\n", ok)
	}
}

func stopTimer() {
	//等待定时器fire后，才调用stop，这时stop返回false
	time.Sleep(2 * timeCount)
	if !myTimer.Stop() {
		fmt.Printf("stopTimer stop Timer get false,try to drain timer's channel\n")
		//让readTimer执行
		time.Sleep(2 * timeCount)
		//这里会死锁，永远卡住
		<-myTimer.C
		fmt.Printf("stopTimer stop Timer success drain timer's channel\n")
	}
}
