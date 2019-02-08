package persist

import (
	"testing"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"fmt"
	"encoding/json"
	"go_study/crawler/model"
)

func TestItemSaver(t *testing.T) {

	expected := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "5000-8000",
		Gender:     "女",
		Xingzuo:    "白羊座",
		Occupation: "人事/行政",
		Marriage:   "未婚",
		House:      "已购房",
		Hukou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
	}

	/*
	//存储数据到ES中

	esId,err := Save(expected)
	if err!=nil{
		panic(err)
	}
	*/

	esId := "AWjMngKUnKIAbCn7xlh7"
	//从ES中取数据
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.227.134:9200"))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(esId).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resp.Source)

	var profile model.Profile
	json.Unmarshal(([]byte)(*resp.Source), &profile)

	if profile == expected {
		fmt.Println("itemSaver function test Passed")
	} else {
		fmt.Println("itemSaver function test failed")
	}
}
