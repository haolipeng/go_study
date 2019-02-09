package persist_distributed

import (
	"go_study/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"go_study/crawler/engine"
)

type ItemSaverService struct {
	client  *elastic.Client
	esIndex string
}

//返回保存是否成功,传指针
func (service *ItemSaverService) Save(item engine.Item, result *string) error {

	err := persist.Save(service.client, service.esIndex, item)
	if err != nil {
		*result = "failed"
	} else {
		*result = "success"
	}

	return err
}
