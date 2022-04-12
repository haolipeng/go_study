package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonRawMessage(t *testing.T) {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)

	//json反序列化
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		panic(0)
	}

	//遍历切片，根据Space字段来解析不同类型数据
	for _, color := range colors {
		var dst interface{}
		switch color.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		//反序列化RawMessage到对象
		err = json.Unmarshal(color.Point, dst)
		if err != nil {
			fmt.Println("json.Unmarshal RawMessage failed")
		}
		fmt.Println(color.Space, dst)
	}
}
