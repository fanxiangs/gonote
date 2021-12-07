package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createChannel() {
	fmt.Printf(" run in creat goroutine\n")
	go func() {
		fmt.Printf(" run in child goroutine\n")
		go func() {
			fmt.Printf(" run in grand child goroutine\n")
		}()
	}()
}

func readChannel() {
	var ch chan int = make(chan int, 4)
	for i := 0; i < cap(ch); i++ {
		ch <- i // 写通道
	}
	for len(ch) > 0 {
		var value int = <-ch // 读通道
		fmt.Println(value)
	}

}

func send(ch chan int) {
	for i := 0; i < 5; i++ {
		var value = rand.Intn(10)
		ch <- value
		fmt.Printf("send %d\n", value)
	}
}
func rec(ch chan int) {
	for {
		value := <-ch
		fmt.Printf("rec %d\n", value)
		time.Sleep(time.Second)
	}
}
func blockChannel() {
	var ch = make(chan int, 0)
	go rec(ch)
	send(ch)
	close(ch)
}

// 没隔一会生产一个数
func sendT(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一的目标通道
func collect(source chan int, target chan int) {
	for v := range source {
		target <- v
	}
}

func recT(ch chan int) {
	for v := range ch {
		fmt.Printf("receive %d\n", v)
	}
}

func multipleChannel() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go sendT(ch1, time.Second)
	go sendT(ch2, 2*time.Second)
	go collect(ch1, ch3)
	go collect(ch2, ch3)
	recT(ch3)
}

func selectSend(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}
func selectRec(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("recv %d from ch1\n", v)
		case v := <-ch2:
			fmt.Printf("recv %d from ch2\n", v)
		}
	}
}

func selectChannel() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go selectSend(ch1, time.Second)
	go selectSend(ch2, 2*time.Second)
	selectRec(ch1, ch2)
}
func noSend(ch1 chan int, ch2 chan int) {
	i := 0
	for {
		select {
		case ch1 <- i:
			fmt.Printf("send ch1 %d\n", i)
		case ch2 <- i:
			fmt.Printf("send ch2 %d\n", i)
		default:
		}
	}
}

func noRec(ch chan int, gap time.Duration, name string) {
	for v := range ch {
		fmt.Printf("receive %s %d\n", name, v)
		time.Sleep(gap)
	}
}

func noBlockChannel() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go noRec(ch1, time.Second, "ch1")
	go noRec(ch2, 2*time.Second, "ch2")
	noSend(ch1, ch2)
}
func main() {
	//createChannel()
	//time.Sleep(time.Second)
	//fmt.Println("main goroutine will quit")
	//readChannel()		// 读写通道
	//blockChannel()	// 读写堵塞
	//multipleChannel() // 多路通道
	//selectChannel()
	noBlockChannel() // 非阻塞读写
}
