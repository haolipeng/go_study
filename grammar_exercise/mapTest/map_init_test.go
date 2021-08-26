package mapTest

import (
	"fmt"
	"testing"
)

type VulnInfos map[string]string

func TestMapInit(t *testing.T) {
	relation := VulnInfos{}
	relation["first"] = "haolipeng"
	relation["second"] = "zhouyang"
	fmt.Println(relation)

	var nilMap map[string]string
	if nilMap == nil {
		fmt.Println("map not init")
	}
	nilMap = make(map[string]string)
	fmt.Println("after make function:")
	if nilMap == nil {
		fmt.Println("map not init")
	} else {
		fmt.Println("map already init")
	}
	//nilMap["name"] = "haolipeng" //panic: assignment to entry in nil mapTest
}
