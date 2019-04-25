package engine

type ConcurrentEngine struct {
	//并发引擎中有调度器来源
	Scheduler Scheduler

	WorkerCount int

	ItemChan chan Item
}

type Scheduler interface {
	//组合方式
	ReadyNotifier

	// interface{}定义方法不需要参数名 指明参数类型即可
	//提交任务
	Submit(Request)
	//会改变
	// 把Run方法生成的 in := make(chan Request) 放到接口中
	//ConfigureMasterWorkerChan(chan Request)

	//我有一个worker，请问给我哪个chan
	WorkerChan() chan Request

	Run()
}

//把scheduler整个送给createworker有点重,所以拆分出来一个接口
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 所有worker共用1个input/output channel
	//创建worker
	//in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//找scheduler要chan
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//创建worker后再submit
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	//收取out的结果
	// 程序没有退出条件 一直轮询等待新数据
	for {
		result := <-out

		//打印item
		for _, item := range result.Items {
			//log.Printf("Got items #%d:%v", itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}

		//把request再送给调度器scheduler
		for _, request := range result.Requests {
			// 这里结构体值传递
			//把新的request传进去

			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
//每个worker对外接口是chan request
func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			//从in拿数据
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
