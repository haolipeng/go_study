package mapTest

import (
	"fmt"
	"testing"
)

//ServiceInfo 服务信息
type ServiceInfo struct {
	Ip   string
	Port int
}

func TestMapSliceAppend(t *testing.T) {
	var (
		infos []ServiceInfo
	)
	//以protocol协议类型作为key
	relation := make(map[string][]ServiceInfo)

	//添加一条记录
	si := ServiceInfo{
		Ip:   "127.0.0.1",
		Port: 80,
	}
	infos = append(infos, si)

	relation["first"] = infos

	//在原有记录的基础上追加记录
	if v, ok := relation["first"]; ok {
		v = append(v, ServiceInfo{
			Ip:   "192.168.1.1",
			Port: 21,
		})
	}

	//打印map，查看追加的记录是否
	fmt.Println(relation) //映射表中只有Ip:   "127.0.0.1", Port: 80,选项
}
