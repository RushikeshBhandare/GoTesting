package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	counter := 0
	const GRR = 100

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(GRR)
	for i := 0; i < GRR; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Num of CPU :", runtime.NumCPU())
	fmt.Println("Num of GoRoutune:", runtime.NumGoroutine())

	fmt.Println("counter :", counter)
}
