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

func (p Packet) Encode() []byte {
	//内容长度
	buf2 := new(bytes.Buffer)
	length := len(p.info)
	err := binary.Write(buf2, binary.LittleEndian, (int32)(length)) // json数据长度
	checkError(err)

	//内容本身
	err = binary.Write(buf2, binary.LittleEndian, []byte(p.info)) // json数据
	checkError(err)

	buf := &bytes.Buffer{}
	p.length = uint32(buf2.Len() + 8)
	err = binary.Write(buf, binary.LittleEndian, p.length) // json数据长度+8校验码长度 总长度
	checkError(err)

	p.crc32 = crc32.ChecksumIEEE(buf2.Bytes())
	err = binary.Write(buf, binary.LittleEndian, p.crc32) // 校验码
	checkError(err)

	err = binary.Write(buf, binary.LittleEndian, buf2.Bytes()) // json数据信息
	checkError(err)
	return buf.Bytes()
}

func (p *Packet) Decode(buff []byte) {
	//fmt.Println("Whole buffer:", buff)
	buf := bytes.NewBuffer(buff)
	err := binary.Read(buf, binary.LittleEndian, &(p.length))
	checkError(err)
	fmt.Println("The total length: ", p.length)

	err = binary.Read(buf, binary.LittleEndian, &(p.crc32))
	checkError(err)

	buf2 := bytes.NewBuffer(buff[8:])
	crc := crc32.ChecksumIEEE(buf2.Bytes())
	if crc != p.crc32 {
		fmt.Println(" crc not check")
	}

	p.info = (string)(buf2.Bytes())
	fmt.Printf("origin json string:%s\n", p.info)
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
