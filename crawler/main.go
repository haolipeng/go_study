package main

import (
	"go_study/crawler/engine"
	"go_study/crawler/zhenai/parser"
	"go_study/crawler/scheduler"
)

const seed = "http://www.zhenai.com/zhenghun"

func main() {
	//simple engine version code
	/*
	e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:seed,
		ParserFunc:parser.ParseCityList,
	})
	*/

	//concurrent engine version code
	e := engine.ConcurrentEngine{&scheduler.QueuedScheduler{}, 10}
	e.Run(engine.Request{
		Url:        seed,
		ParserFunc: parser.ParseCityList,
	})
}
