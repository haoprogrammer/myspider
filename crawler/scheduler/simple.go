package scheduler

import "haoprogrammer/myspider/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {

	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//将请求发送到worker
	//每个request都会建一个goroutine,往统一的workerChan里面分发
	go func() { s.workerChan <- r }()

}
