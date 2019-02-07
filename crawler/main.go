package main

import (
	"go_study/crawler/engine"
	"go_study/crawler/scheduler"
	"go_study/crawler/zhenai/parser"
	"go_study/crawler/persist"
)

const seed = "http://www.zhenai.com/zhenghun"
const shanghai_seed = "http://www.zhenai.com/zhenghun/shanghai"

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
	/*e := engine.ConcurrentEngine{&scheduler.QueuedScheduler{}, 10}
	e.Run(engine.Request{
		Url:        seed,
		ParserFunc: parser.ParseCityList,
	})
	*/

	//parser target city
	e := engine.ConcurrentEngine{
		&scheduler.QueuedScheduler{},
		10,
		persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        shanghai_seed,
		ParserFunc: parser.ParserCity,
	})
}
