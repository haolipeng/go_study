package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

type packet struct {
	Sensid uint16
	Locid  uint16
	Ts     uint32
	Temp   uint16
}

func TestReadToStruct(t *testing.T) {
	//1.结构体编码为buf缓冲区
	dataIn := packet{
		Sensid: 1,
		Locid:  1233,
		Ts:     123452123,
		Temp:   12,
	}

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, dataIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	//2.将buf缓冲区内容解码为
	var dataOut packet
	err = binary.Read(buf, binary.BigEndian, &dataOut)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dataOut)
}
