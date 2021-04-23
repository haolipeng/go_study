package main

import "fmt"

// 定义一个接口，里面有两个方法
type pluginfunc interface {
	hello()
}

// 定义一个类，来存放我们的插件
type plugins struct {
	plist map[string]pluginfunc
}

// 初始化插件
func (p *plugins) init() {
	p.plist = make(map[string]pluginfunc)
}

// 注册插件到映射表中
func (p *plugins) register(name string, plugin pluginfunc) {
	p.plist[name] = plugin
}

//////////////////////////////////////////////////////
//plugin1
type plugin1 struct{}

func (p *plugin1) hello() {
	fmt.Println("plugin1 hello")
}

//plugin2
type plugin2 struct{}

func (p *plugin2) hello() {
	fmt.Println("plugin2 hello")
}

//plugin3
type plugin3 struct{}

func (p *plugin3) hello() {
	fmt.Println("plugin3 hello")
}

//////////////////////////////////////////////////////

func main() {
	pluginManager := new(plugins)
	pluginManager.init()

	//申请插件对象
	p1 := new(plugin1)
	p2 := new(plugin2)
	p3 := new(plugin3)

	//注册插件
	pluginManager.register("plugin1", p1)
	pluginManager.register("plugin2", p2)
	pluginManager.register("plugin3", p3)

	for _, v := range pluginManager.plist {
		v.hello()
	}
}
