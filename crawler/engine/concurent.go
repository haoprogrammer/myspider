package engine

import (
	"log"
)

type ConcurrentEngine struct {
	//并发引擎中有调度器来源
	Scheduler Scheduler

	WorkerCount int
}

type Scheduler interface {
	// interface{}定义方法不需要参数名 指明参数类型即可
	//提交任务
	Submit(Request)

	//会改变
	// 把Run方法生成的 in := make(chan Request) 放到接口中
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 所有worker共用1个input/output channel
	//创建worker
	in := make(chan Request)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	out := make(chan ParserResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//创建worker后再submit
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	//收取out的结果
	// 程序没有退出条件 一直轮询等待新数据
	for {
		result := <-out

		//打印item
		for _, item := range result.Items {
			log.Printf("Got items #%d:%v", itemCount, item)
			itemCount++
		}

		//把request再送给调度器scheduler
		for _, request := range result.Requests {
			// 这里结构体值传递
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			//tell scheduler i'm ready
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}
