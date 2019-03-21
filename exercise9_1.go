package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go boo()
	fmt.Println("Number of Go routines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func boo() {
	fmt.Println("Hmm..")
	wg.Done()
}
