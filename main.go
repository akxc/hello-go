package main

import (
	"fmt"
	"hello-go/monitor"
	"plugin"
	"time"
)

func package_test() {
	fmt.Printf("hello, world\n")
	monitor.Connect()
}
func variable_visibility() {
	// Jobb is in test.go which is at the same directory of main.go
	// When I run F5 to start debug in IDE, it work fine.
	// But an error ".\main.go:13:9: undefined: Jobb" is reporter with
	// CLI "go run main.go".
	//job := Jobb{"job in main"}
	//fmt.Println("success in IDE, failed with CLI: ", job)
}
func channel_test() {
	//var x chan int
	x := make(chan int, 2)
	go func() {
		fmt.Printf("hello, 1\n")
		x <- 1
		fmt.Printf("hello, 2\n")
		//x <- 2
		//x <- 3
	}()
	fmt.Printf("helloA, %d\n", <-x)
	<-x
	fmt.Printf("helloA, %d\n", <-x)
}
func timeout_by_slelect() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}
func check_chan_full_by_select() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full!")
	}
}
func writeOnlyCh(ch chan<- int) {
	//a := <-ch  //ERROR: receive from readonly type chan<- int
	ch <- 1
	ch <- 1
	ch <- 1
}
func readOnlyCh() <-chan string {
	readCh := make(chan string, 3)
	// wite 3 element to readCh, and read from readCh after call "readCh := readOnlyCh()" will
	// result in "fatal error: all goroutines are asleep - deadlock!"
	// output := []string{"result", "read", "in", "foo"}
	// for _, v := range output {
	// 	readCh <- v
	// }
	readCh <- "Written by readOnlyCh func"
	return readCh
}
func inputWriteOutputRead(writeOnlyCh chan<- int) <-chan string {
	writeOnlyCh <- 2
	readCh := make(chan string, 3)
	readCh <- "Written by inputWriteOutputRead func"
	writeOnlyCh <- 3
	return readCh
}

func one_direction_channel() {
	// One direction channel means only read or write channel.
	// Used mainly for function declaration.
	// func foo(ch chan<- int) <-chan int {...}
	//convention over configuration?

	writeCh := make(chan int)
	//	writeOnlyCh(writeCh) //without a goroutine will result in "fatal error: all goroutines are asleep - deadlock!"
	go writeOnlyCh(writeCh)
	fmt.Println("read from input parameter of writeOnlyCh(ch chan<- int){}:", <-writeCh)
	// <-writeCh //read once more also result in "fatal error: all goroutines are asleep - deadlock!"
	readCh := readOnlyCh()
	fmt.Println("read from output parameter of readOnlyCh() <-chan string{}:", <-readCh)
	// <-readCh //read once more also result in "fatal error: all goroutines are asleep - deadlock!"

	writeChBuf := make(chan int, 2)
	readCh2 := inputWriteOutputRead(writeChBuf)
	fmt.Println("read from output parameter of inputWriteOutputRead() <-chan string{}:", <-readCh2)

	// iterate with range result in "fatal error: all goroutines are asleep - deadlock!"
	// for x := range writeChBuf {
	// 	fmt.Println(x)
	// }
	for {
		select {
		case x := <-writeChBuf:
			fmt.Println("read from writeChBuf: ", x)
		default:
			fmt.Println("writeChBuf is empty!")
			time.Sleep(time.Second * 3)
		}
	}
	//readCh2 <- "try to write "

}

func load_plugin() {
	pdll, err := plugin.Open("print.so")
	if err != nil {
		fmt.Println("err happened at Open plugin: ", err)
		return
	}
	funcPrint, err := pdll.Lookup("PrintByPlugin")
	if err != nil {
		fmt.Println("err happened at Lookup plugin: ", err)
		return
	}
	funcPrint.(func(string))("hello go plugin")
	return
}
func main() {
	//defer_test()
	//package_test()
	//variable_visibility()
	//channel_test()
	//timeout_by_slelect()
	//check_chan_full_by_select()
	//one_direction_channel()
	load_plugin()

}

func defer_test() {
	defer func() { fmt.Println("print start") }()
	defer func() { fmt.Println("print middle") }()
	defer func() { fmt.Println("print end") }()
	panic("trigger exception")
}
