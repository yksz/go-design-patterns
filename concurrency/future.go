package main

import (
	"log"
	"time"
)

type Future chan string

type Executor struct{}

func (e *Executor) Submit(fn func() string) Future {
	future := make(Future)
	go func() {
		val := fn()
		future <- val
	}()
	return future
}

func main() {
	task := func() string {
		log.Println("Task: Now processing...")
		time.Sleep(2 * time.Second)
		log.Println("Task: Process is completed.")
		return "Task has finished!"
	}
	executor := new(Executor)
	future := executor.Submit(task)
	log.Printf("Future result: %s\n", <-future)
}
