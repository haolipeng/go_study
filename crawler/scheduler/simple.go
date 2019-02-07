package scheduler

import "go_study/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

//为每个request请求建立一个协程
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(ch chan engine.Request) {
	s.workerChan = ch
}
