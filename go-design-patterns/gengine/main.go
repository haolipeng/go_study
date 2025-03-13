package main

import (
	"fmt"
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

// Hello报文结构 - 使用string类型
type HelloPacket struct {
	NetworkMask string `json:"network_mask"`
}

// OSPF数据包结构 - 使用string类型
type OSPFPacket struct {
	SrcRouter string       `json:"srcrouter"`
	Hello     *HelloPacket `json:"hello"`
}

// 保持原始规则格式不变
const rule = `
rule "check_ospf" "检查OSPF报文"
begin
    if ((ospf.srcrouter == "192.168.170.3")) && (ospf.hello.network_mask == "255.255.255.0") {
        println("规则匹配成功!")
    }
end
`

func main() {
	// 创建数据上下文
	dataContext := context.NewDataContext()

	// 创建规则构建器
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	// 加载规则
	err := ruleBuilder.BuildRuleFromString(rule)
	if err != nil {
		fmt.Printf("构建规则失败: %v\n", err)
		return
	}

	// 创建规则引擎
	eng := engine.NewGengine()

	// 准备测试数据 - 使用字符串类型的IP地址
	packet := &OSPFPacket{
		SrcRouter: "192.168.170.3", // 直接使用字符串
		Hello: &HelloPacket{
			NetworkMask: "255.255.255.0", // 直接使用字符串
		},
	}

	// 注入数据到上下文
	dataContext.Add("ospf", packet)

	// 执行规则
	err = eng.Execute(ruleBuilder, true)
	if err != nil {
		fmt.Printf("执行规则失败: %v\n", err)
		return
	}
}
