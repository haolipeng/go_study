package parser

import (
	"go_study/crawler/engine"
	"regexp"
)

const cityNextPage = `href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`
//过滤城市列表的正则表达式
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, name)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParserProfile(c, name)
			},
		})

	}

	//获取下一页内容
	re = regexp.MustCompile(cityNextPage)
	matches = re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}
