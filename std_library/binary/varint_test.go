package binary

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestVarint(t *testing.T) {
	inputs := [][]byte{
		{0x81, 0x01},
		{0x7f},
		{0x03},
		{0x01},
		{0x00},
		{0x02},
		{0x04},
		{0x7e},
		{0x80, 0x01},
	}
	for _, b := range inputs {
		x, n := binary.Varint(b)
		if n != len(b) {
			fmt.Println("Varint did not consume all of in")
		}
		fmt.Println(x)
	}
}
