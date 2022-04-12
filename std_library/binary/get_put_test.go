package binary

import (
	"encoding/binary"
	"fmt"
	"testing"
)

//计算机内部处理是小端字节序，因为电路先处理低位字节，效率比较高，计算是从低位开始的。
//人类习惯读写大端字节序，除计算机内部处理采用小端字节序，其他场合都采用大端字节序，如网络传输和文件存储

//实验：小端方式读取值  vs  大端方式读取值
func TestGet(t *testing.T) {
	b := []byte{0xe8, 0x03, 0xd0, 0x07}
	x1 := binary.LittleEndian.Uint16(b[0:])
	x2 := binary.LittleEndian.Uint16(b[2:])
	fmt.Printf("%#06x %#04x %#x\n", x1, x2, x1)

	x1 = binary.BigEndian.Uint16(b[0:])
	x2 = binary.BigEndian.Uint16(b[2:])
	fmt.Printf("%#06x %#04x %#x\n", x1, x2, x1)
}

//实验：小端方式写入值  vs  大端方式写入值
func TestPut(t *testing.T) {
	//字节流原本的样子
	originB := []byte{0xe8, 0x03, 0xd0, 0x07}
	fmt.Printf("origin buffer:% x\n", originB)

	//以小端方式写入值
	buf1 := make([]byte, 16)
	binary.LittleEndian.PutUint16(buf1[0:], 0x03e8)
	binary.LittleEndian.PutUint16(buf1[2:], 0x07d0)
	fmt.Printf("little endian buffer:%x\n", buf1) //和原字符串相符

	//以大端方式写入值
	buf2 := make([]byte, 16)
	binary.BigEndian.PutUint64(buf2[0:], 0x03e8)
	binary.BigEndian.PutUint64(buf2[8:], 0x07d0)
	fmt.Printf("big endian buffer:%x\n", buf2)
}
