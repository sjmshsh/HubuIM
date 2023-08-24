package tcp

import (
	"bytes"
	"encoding/binary"
)

type DataPgk struct {
	Len  uint32
	Data []byte
}

func (d *DataPgk) Marshal() []byte {
	// 创建了一个新的字节缓冲区, 用于将数据写入其中.
	// 初始时, 缓冲区是空的，因此传递一个空的字节数组[]bytes{}给构造函数
	bytesBuffer := bytes.NewBuffer([]byte{})
	// 将d.Len字段按照大端序的方式写入到字节缓冲区中, 以二进制格式
	binary.Write(bytesBuffer, binary.BigEndian, d.Len)
	// 获取字节缓冲区中的数据作为字节数组, 并使用append将d.Data字段数据追加到其中
	return append(bytesBuffer.Bytes(), d.Data...)
}
