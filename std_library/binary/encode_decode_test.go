package binary

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"
	"testing"
)

//暂时没有掌握好
type Register struct {
	ACTION string
	SID    int32
}

type Packet struct {
	length uint32
	crc32  uint32
	info   string
}

const CrcLength = 8

/*
0   1   2   3   4   5   6   7   8   9   10  11  12  13  14 (bytes)
+-------------------------------+---------------+-------+
|       ID                 		|Timestamp           | Value |
+-------+-----------------------+---------------+-------+
*/
func (p Packet) Encode() []byte {
	//info内容长度
	innerBuf := new(bytes.Buffer)
	var length int = len([]byte(p.info))
	err := binary.Write(innerBuf, binary.LittleEndian, (int32)(length)) // json数据长度
	checkError(err)
	fmt.Println("encode, inner buf length:", length)

	//写入info内容到innerBuf中
	err = binary.Write(innerBuf, binary.LittleEndian, []byte(p.info)) // json数据
	checkError(err)
	fmt.Println("encode, inner buf:", p.info)

	//总长度
	buf := &bytes.Buffer{}
	p.length = uint32(innerBuf.Len() + 8) //偏移length +crc32 = 8个字节
	err = binary.Write(buf, binary.LittleEndian, p.length)
	checkError(err)
	fmt.Println("encode, p.length:", p.length)

	//计算内部缓存innerBuf的crc32
	p.crc32 = crc32.ChecksumIEEE(innerBuf.Bytes())
	err = binary.Write(buf, binary.LittleEndian, p.crc32) // 校验码
	checkError(err)
	fmt.Println("encode, p.crc32:", p.crc32)

	//写入json数据
	err = binary.Write(buf, binary.LittleEndian, innerBuf.Bytes()) // json数据信息
	checkError(err)
	return buf.Bytes()
}

func (p *Packet) Decode(buff []byte) {
	buf := bytes.NewBuffer(buff)
	err := binary.Read(buf, binary.LittleEndian, &(p.length))
	checkError(err)
	fmt.Println("decode, p.length: ", p.length)

	err = binary.Read(buf, binary.LittleEndian, &(p.crc32))
	checkError(err)
	fmt.Println("decode, p.crc32:", p.crc32)

	innerBuf := bytes.NewBuffer(buff[8:]) //偏移length +crc32 = 8个字节
	crc := crc32.ChecksumIEEE(innerBuf.Bytes())
	if crc != p.crc32 {
		fmt.Println(" crc not check")
	}

	var length int32
	err = binary.Read(innerBuf, binary.LittleEndian, &length)
	checkError(err)

	err = binary.Read(innerBuf, binary.LittleEndian, p.info)
	fmt.Printf("decode, p.info length:%s\n", p.info)

	p.info = (string)(innerBuf.Bytes())
	fmt.Printf("decode, p.info content:%s\n", p.info)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func TestEncodeDecode(t *testing.T) {
	m := Register{"asdjfiosdjfoijsdoifjsaodfjosiadjfiosdjfiosdjfoisdjfoisjadiofjsdoijfsoidfjsoidfjsddfjslkdjfksdjfkdlsjfklsdjfkljsdklfjsdlkfjoiwrejiojsiodjfiofffffffffffffsdklfjoirejufoisjdfoiajfoisjdfoiajsiofsjdoifjosaidjfsaijdfiosapjfoisjdoifjodsijfosidjfoisdjfosidfjsdoifjsidojfosidjfsdoioi", 6}
	b, err := json.Marshal(m)
	checkError(err)

	var packet Packet
	packet.info = string(b)
	buf := packet.Encode()
	fmt.Println(len(buf))

	var msg Packet
	msg.Decode(buf)
}
