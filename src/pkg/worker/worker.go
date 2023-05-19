package worker

import "sync"

type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		task()
	}
}

func (p *WorkerPool) Add(task func()) {
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	close(p.tasks)
	p.wg.Wait()
}

func NewWorkerPool(numWorkers uint) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func()),
	}

	pool.wg.Add(int(numWorkers))

	for i := 0; i < int(numWorkers); i++ {
		go pool.worker()
	}

	return pool
}
