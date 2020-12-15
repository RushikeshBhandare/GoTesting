package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	var counter int64
	const GRR = 100

	var wg sync.WaitGroup
	wg.Add(GRR)
	for i := 0; i < GRR; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	fmt.Println("counter :", counter)
}
