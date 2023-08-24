package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func ReadData(conn *net.TCPConn) ([]byte, error) {
	var dataLen uint32
	dataLenBuf := make([]byte, 4)
	if err := readFixedData(conn, dataLenBuf); err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(dataLenBuf)
	if err := binary.Read(buffer, binary.BigEndian, &dataLen); err != nil {
		return nil, fmt.Errorf("read headlen error:%s", err.Error())
	}
	if dataLen <= 0 {
		return nil, fmt.Errorf("wrong headLen : %d", dataLen)
	}
	dataBuf := make([]byte, dataLen)
	if err := readFixedData(conn, dataBuf); err != nil {
		return nil, fmt.Errorf("read headlen error:%s", err.Error())
	}
	return dataBuf, nil
}

// 读取固定buf长度的数据
func readFixedData(conn *net.TCPConn, buf []byte) error {
	// 设置了读取操作的截止时间，当前读取最多会等待120s
	_ = (*conn).SetReadDeadline(time.Now().Add(time.Duration(120) * time.Second))
	var pos int = 0
	var totalSize = len(buf)
	for {
		c, err := (*conn).Read(buf[pos:])
		if err != nil {
			return err
		}
		pos = pos + c
		if pos == totalSize {
			break
		}
	}
	return nil
}
