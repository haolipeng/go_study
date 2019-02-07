package engine

import (
	"go_study/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(seeds ...Request) {
	//将seeds中元素导入切片requests中
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//requests队列起缓存元素
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests ...)

		//打印items
		for _, item := range parserResult.Items {
			log.Printf("Get Item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.FetchWithUserAgent(r.Url)

	if err != nil {
		log.Printf("Fetcher:error "+
			"fetching url:%s %v", r.Url, err)
		return ParseResult{}, err
	}

	parserResult := r.ParserFunc(body)

	return parserResult, nil
}
