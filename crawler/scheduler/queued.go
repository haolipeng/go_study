package scheduler

import "go_study/crawler/engine"

//requestChan和workerChan是何时创建的
type QueuedScheduler struct {
	RequestChan chan engine.Request      //chan是一个worker类型，每个worker有自己的channel
	WorkerChan  chan chan engine.Request //存储worker队列的队列
}

func (s *QueuedScheduler) GetWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(ch chan engine.Request) {

}

//有一个worker已经准备好了，可以接收Request请求了
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

//创建worker channel和request channel,
func (s *QueuedScheduler) Run() {
	//创建workerChan和requestChan
	// Engine引擎向request队列中添加元素
	// 在创建worker工作者时，将worker对应的队列反注册给QueuedScheduler
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ [] chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			//request队列和worker队列中数量大于0
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]

				//****此处如果阻塞，后续将无法收到requestChan或workerChan上的请求了
			}

			//request和worker事件的发生是彼此独立的
			select {
			case activeWorker <- activeRequest:
				//从队列中弹出元素
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			case rch := <-s.RequestChan:
				requestQ = append(requestQ, rch)
			case wch := <-s.WorkerChan:
				workerQ = append(workerQ, wch)
			}
		}
	}()
}
