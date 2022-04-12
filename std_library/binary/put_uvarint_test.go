package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

//TODO:这个题暂时没搞清楚
func TestPutUVarint(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Println("put bytes:", n)
		fmt.Printf("%x\n", buf[:n])
		fmt.Printf("buf: %x\n", buf)
		fmt.Println("-------------------------------")
	}
}

func TestRead(t *testing.T) {
	var pi float64

	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.LittleEndian, &pi); err != nil {
		return
	}

	fmt.Println(pi)
}
