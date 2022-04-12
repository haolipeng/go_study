package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

/*
brief:将PI值写入buf缓冲区中
一个字节是8bit，也就是8位，整数最大表示是255
18 2d 44 54 fb 21 09 40  十六进制整数表示
[24 45 68 84 251 33 9 64]
用8byte 表示了 math.Pi

Write: 将data 转为bytes
*/
func TestWrite(t *testing.T) {
	var buf bytes.Buffer
	var pi float64 = math.Pi
	err := binary.Write(&buf, binary.LittleEndian, pi)
	if err != nil {
		t.Logf("binary.Write failed")
	}
	fmt.Printf("%x\n", buf.Bytes())
	fmt.Println(buf.Bytes())
}
