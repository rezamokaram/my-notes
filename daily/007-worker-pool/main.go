package main

import (
	"fmt"
	"sync"
)

type Task func()

type ThreadPool struct {
	tasks    chan Task
	wg       sync.WaitGroup
	stopChan chan struct{}
}

func NewThreadPool(numWorkers int) *ThreadPool {
	pool := &ThreadPool{
		tasks:    make(chan Task),
		stopChan: make(chan struct{}),
	}

	for i := 0; i < numWorkers; i++ {
		go pool.worker(i)
	}

	return pool
}

func (p *ThreadPool) worker(workerID int) {
	for {
		select {
		case task := <-p.tasks:
			fmt.Printf("Worker %d is executing task...\n", workerID)
			task()
			p.wg.Done()
		case <-p.stopChan:
			fmt.Printf("Worker %d stopping.\n", workerID)
			return
		}
	}
}

func (p *ThreadPool) Submit(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *ThreadPool) Wait() {
	p.wg.Wait()
}

func (p *ThreadPool) Stop() {
	close(p.stopChan)
	close(p.tasks)
}

func main() {
	pool := NewThreadPool(3)

	for i := 0; i < 5; i++ {
		taskID := i
		pool.Submit(func() {
			fmt.Printf("Task %d is running\n", taskID)
		})
	}
	pool.Wait()


	// for i := 0; i < 5; i++ {
	// 	taskID := i
	// 	pool.Submit(func() {
	// 		fmt.Printf("Task %d is running\n", taskID)
	// 	})
	// }
	// pool.Wait()


	pool.Stop()
}
