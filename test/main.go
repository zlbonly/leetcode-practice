package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//printGoroutineOrder()
	timeOut()
}

func timeOut() {
	ch1 := make(chan struct{})
	go func() {
		fmt.Printf("%v\n", time.Now().Unix())
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}()
	ch2 := make(chan struct{})

	select {
	case _, ok := <-ch2:
		if ok {
			fmt.Printf("%v\n", "zlbtitie")
		}
	case _, ok := <-ch1:
		if ok {
			fmt.Printf("%v\n", time.Now().Unix())

			fmt.Printf("%v\n", "time out")
		}
	}
	close(ch1)
	close(ch2)
}

func printGoroutineOrder() {
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		if _, ok := <-ch1; ok {
			fmt.Printf("执行%v\n", "A")
			ch2 <- true
		}
	}()

	go func() {
		defer wg.Done()
		if _, ok := <-ch2; ok {
			fmt.Printf("执行%v\n", "B")
			ch3 <- true
		}
	}()
	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
			wg.Done()
		}()
		if _, ok := <-ch3; ok {
			fmt.Printf("执行%v\n", "C")
			//ch2<-true
		}
	}()
	ch1 <- true
	wg.Wait()
}

//
//A首先被a阻塞，A()结束后关闭b，使b可读
func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

// B首先被a阻塞，B()结束后关闭b，使b可读
func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

// C首先被a阻塞
func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}
