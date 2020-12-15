package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("Oprating System ", runtime.GOOS)

	wg.Add(1)
	go one()
	two()

	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("Goroutine", runtime.NumGoroutine())
	fmt.Println("Oprating System ", runtime.GOOS)
	wg.Wait()
}

func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("One : ", i)
	}
	wg.Done()
}

func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("two : ", i)
	}
}
