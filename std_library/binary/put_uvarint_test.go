package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestUVarint(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
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
