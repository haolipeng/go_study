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

func main() {
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
