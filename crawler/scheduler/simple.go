package scheduler

import "haoprogrammer/myspider/crawler/engine"

//所有的worker共用一个chan
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	//所有的worker共用同一个worker
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//将请求发送到worker
	//每个request都会建一个goroutine,往统一的workerChan里面分发
	go func() { s.workerChan <- r }()

}
