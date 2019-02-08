package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"fmt"
	"context"
)

//获取itemSaver的input 通道
func ItemSaver() chan interface{} {
	out := make(chan interface{})

	//print item
	go func() {
		itemCount := 0
		for ; ; {
			item := <-out
			log.Printf("Item Saver: Got item "+
				"#%d: %v", itemCount, item)

			itemCount++
		}
	}()

	return out
}

func Save(item interface{}) (id string, err error) {
	//1.创建elasticSearch client
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.227.134:9200"))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v", resp)

	return resp.Id, nil
}