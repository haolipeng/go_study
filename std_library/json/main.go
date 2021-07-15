package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

//重命名序列化后的名字
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

//验证序列化和反序列化所使用的结构体不一致是否可行
type responseExpand struct {
	Page    int
	Fruits  []string
	Count   int64
	Content map[string]string
}

type Task struct {
	//TODO:任务结构体各字段对应的json字符串需和前端联调
	//检测目标（有不同的主机分组）
	TargetHosts map[string][]string `json:"target_hosts"` //key = 业务组,value = 业务组下的主机列表

	//检测项（如系统风险、应用风险、账号风险、容器风险）
	DetectItems []string `json:"detect_items"`

	//是否获取操作系统
	ObtainOS bool `json:"obtain_os"`

	//检测网段
	NetworkSegment string `json:"network_segment"`

	//任务类型（即时任务 or 定时任务）
	TaskType   int `json:"task_type"`   //任务类型
	TaskPeriod int `json:"task_period"` //任务下发间隔
}

func jsonTask2String() {
	targetHost := make(map[string][]string)
	targetHost["业务组1"] = []string{
		"192.168.100.1",
		"192.168.100.2",
	}
	targetHost["业务组2"] = []string{
		"192.168.200.3",
		"192.168.200.4",
	}
	targetHost["业务组3"] = []string{
		"192.168.222.5",
		"192.168.222.6",
	}

	detectItems := []string{
		"账号相关文件权限检测",
		"系统内核参数检测校验",
	}

	task := Task{
		TargetHosts:    targetHost,
		DetectItems:    detectItems,
		ObtainOS:       false,
		NetworkSegment: "192.168.1.100",
		TaskType:       0,
		TaskPeriod:     10,
	}

	taskStr, err := json.Marshal(&task)
	if err != nil {

	}
	fmt.Println(taskStr)
}

func main() {
	jsonTask2String()
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	var obj responseExpand
	err := json.Unmarshal(res1B, &obj)
	if err != nil {
		fmt.Println("json.Unmarshal failed")
	}
	fmt.Println(obj)
	//结论：json序列化和反序列化采用的结构体不是同一个，则多余的结构体资源不填写

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err = json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)
}
