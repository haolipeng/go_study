package binary

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestUVarint(t *testing.T) {
	inputs := [][]byte{
		{0x01},
		{0x02},
		{0x7f},
		{0x80, 0x01},
		{0xff, 0x01},
		{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b) // 将 byte 数据 转为对应的 uint64整数
		if n != len(b) {
			fmt.Println("Uvarint did not consume all of in")
		}
		fmt.Println(x)
	}
}
