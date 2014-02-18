package main

import (
	"log"
	"sync"
	"time"
)

const (
	num = 5
)

var (
	busy bool
	m    sync.Mutex
	wg   sync.WaitGroup
)

func execute(id int) {
	defer wg.Done()
	m.Lock()
	if busy {
		log.Printf("Goroutine-%d: Process is busy!\n", id)
		m.Unlock()
		return
	}
	busy = true
	m.Unlock()
	log.Printf("Goroutine-%d: Now processing...\n", id)
	doSomething()
	log.Printf("Goroutine-%d: Process is completed.\n", id)
	busy = false
}

func doSomething() {
	time.Sleep(1 * time.Second)
}

func main() {
	wg.Add(num)
	for i := 0; i < num; i++ {
		go execute(i)
	}
	wg.Wait()
}
