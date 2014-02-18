package main

import (
	"log"
	"time"
)

var (
	stop Runner
)

func init() {
	stop = new(stopRunner)
}

type Runner interface {
	Run()
}

type stopRunner struct{}

func (r *stopRunner) Run() {}

type ThreadPool struct {
	queue     chan Runner
	threads   []*thread
	isRunning bool
}

func NewThreadPool(maxQueueSize int, numberOfThreads int) *ThreadPool {
	queue := make(chan Runner, maxQueueSize)
	threads := make([]*thread, numberOfThreads)
	for i := 0; i < len(threads); i++ {
		threads[i] = newThread(i, queue)
	}
	return &ThreadPool{queue, threads, false}
}

func (p *ThreadPool) Start() {
	if p.isRunning {
		panic("threads have been already started")
	}
	p.isRunning = true
	for _, thread := range p.threads {
		thread.start()
	}
}

func (p *ThreadPool) Stop() {
	if !p.isRunning {
		panic("threads have not been started")
	}
	for i := 0; i < len(p.threads); i++ {
		p.Dispatch(stop)
	}
	p.isRunning = false
	for _, th := range p.threads {
		th.join()
	}
}

func (p *ThreadPool) Dispatch(runner Runner) {
	if !p.isRunning {
		panic("this pool has not been started yet")
	}
	p.queue <- runner
}

type thread struct {
	id    int
	queue chan Runner
	fin   chan bool
}

func newThread(id int, queue chan Runner) *thread {
	return &thread{id, queue, make(chan bool)}
}

func (t *thread) start() {
	log.Printf("Thread-%d start\n", t.id)
	go func() {
		for {
			runner := <-t.queue
			if runner == stop {
				t.fin <- true
				return
			} else {
				runner.Run()
			}
		}
	}()
}

func (t *thread) join() {
	log.Printf("Thread-%d join\n", t.id)
	<-t.fin
}

type ConcreteRunnerA struct{}

func (r *ConcreteRunnerA) Run() {
	time.Sleep(1 * time.Second)
	log.Println("ConcreteRunnerA.Run()")
}

type ConcreteRunnerB struct{}

func (r *ConcreteRunnerB) Run() {
	time.Sleep(2 * time.Second)
	log.Println("ConcreteRunnerB.Run()")
}

func main() {
	runnerA := new(ConcreteRunnerA)
	runnerB := new(ConcreteRunnerB)
	threadPool := NewThreadPool(1, 3)
	threadPool.Start()
	threadPool.Dispatch(runnerA)
	threadPool.Dispatch(runnerA)
	threadPool.Dispatch(runnerB)
	threadPool.Stop()
}
