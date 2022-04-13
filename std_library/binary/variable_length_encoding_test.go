package binary

import (
	"encoding/binary"
	"fmt"
	"testing"
)

type Data struct {
	ID        uint32 //4
	Timestamp uint64 //8
	Value     int16  //2
}

func TestFixedLengthPut(t *testing.T) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, 10)
	fmt.Printf("encoded binary length: %d\n", len(buf)) //占用8个字节
}

func TestVariableLengthPut(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)
	var x uint64 = 10
	n := binary.PutUvarint(buf, x)
	fmt.Printf("encoded binary length: %d\n", len(buf[:n])) //占用1个字节
}

func TestVariableLengthPutAndGet(t *testing.T) {
	data := &Data{
		ID:        100,
		Timestamp: 1600000000,
		Value:     10,
	}

	//正常写入结构体，需要14个字节，用变长方式写入，只需要7个字节
	//变长字段的写入
	buf := make([]byte, 7)

	n := binary.PutUvarint(buf, uint64(data.ID))
	fmt.Println("first n:", n)

	n += binary.PutUvarint(buf[n:], data.Timestamp)
	fmt.Println("second n:", n)

	n += binary.PutVarint(buf[n:], int64(uint64(data.Value)))
	fmt.Println("third n:", n)

	fmt.Printf("encoded binary length: %d\n", len(buf[:n]))

	//解析id
	id, idLen := binary.Uvarint(buf)
	ts, tsLen := binary.Uvarint(buf[idLen:])
	value, _ := binary.Varint(buf[idLen+tsLen:])

	decodedData := &Data{
		ID:        uint32(id),
		Timestamp: uint64(ts),
		Value:     int16(value),
	}
	fmt.Printf("ID: %d, Timestamp: %d, Value: %d\n", decodedData.ID, decodedData.Timestamp, decodedData.Value)
	// ID: 100, Timestamp: 1600000000, Value: 10
}
