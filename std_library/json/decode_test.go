package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Response struct {
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
}

type DataInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonDecoder(t *testing.T) {
	const jsonStream = `
	{
    "data": [
        {
            "name": "haolipeng",
            "age": 32
        },
        {
            "name": "zhouyang",
            "age": 33
        }
    ],
    "message": "success",
    "status": 200
	}`

	var (
		resp  Response
		infos []DataInfo
	)
	//进行json反序列化
	err := json.Unmarshal([]byte(jsonStream), &resp)
	if err != nil {
		fmt.Printf("json.Unmarshal error: %s", err)
		return
	}

	//将resp.Data反序列化成特定类型Info
	err = json.Unmarshal(resp.Data, &infos)
	if err != nil {
		return
	}
	fmt.Println("value:", infos)
}
