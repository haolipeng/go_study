package main

import (
	"go_study/crawler_distributed/persist_distributed"
	"gopkg.in/olivere/elastic.v5"
	"go_study/crawler_distributed/rcpsupport"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	rcpsupport.ServerRpc(":1234", persist_distributed.ItemSaverService{
		client:  client,
		esIndex: "dating_profile",
	})

}
