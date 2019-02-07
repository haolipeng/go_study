package parser

import (
	"go_study/crawler/engine"
	"regexp"
)

//过滤城市列表的正则表达式
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//解析城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//暂时先解析10个城市的数值

	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity, //这里调用ParserCity
		})
	}

	return result
}
