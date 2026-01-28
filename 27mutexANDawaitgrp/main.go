package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - LearnGoWeb.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)

	// Imagine this : SCORE is a shared notebook
	// * Multiple people (goroutines) want to:
	// * Write numbers in it
	// * Read from it
	// Only one writer can write at a time...
	// Readers must not read while someone is writing..
	// That’s exactly what RWMutex enforces...

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

// Why RWMutex instead of Mutex?
// *Lock() → exclusive (writers)
// *RLock() → shared (readers)

// Using RWMutex allows multiple readers to access the shared resource concurrently,

// The code demonstrates how to safely read and write shared data using goroutines, a wait group, and a mutex,
// while showing why synchronization is necessary in concurrent Go programs.

// *3 writers use Lock()
// *1 reader uses RLock()
