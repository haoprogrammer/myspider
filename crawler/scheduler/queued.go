package scheduler

import "haoprogrammer/myspider/crawler/engine"

//每个worker有自己的chan
type QueuedScheduler struct {
	requestChan chan engine.Request
	//workerChan chan worker(chan engine.Request)
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)

}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	//在run里面生成chan
	//改变了s  QueuedScheduler的内容,需要将内容改为指针接受者
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		//存放队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {

			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//队列里面既有request又有worker就可以发
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			//什么时候有一个request和什么时候有一个workready是两个独立的事情
			//让所有chan操作都放到select里面
			select {
			case r := <-s.requestChan:
				//send r to a   ?worker
				//收到消息让其排队
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				//send ?next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}

	}()

}
