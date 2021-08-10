package main

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

	var resp Response
	//进行json反序列化
	err := json.Unmarshal([]byte(jsonStream), &resp)
	if err != nil {
		fmt.Printf("json.Unmarshal error: %s", err)
		return
	}

	/*
		//将resp.Data再次序列化
		bytes, err := json.Marshal(resp.Data)
		if err != nil {
			return
		}

		//将resp.Data反序列化成特定类型Info
		err = json.Unmarshal(bytes, &infos)
		if err != nil {
			return
		}
		fmt.Println("value:", infos)
	*/
}
