package tcp

import "net"

// 将指定数据包发送到网络中
// 这样实现是为了确保完整的发送整个字节数组，即使在一次调用中无法一次性全部发送
// 以下的清空无法一次性将数据发送出去
// 1. TCP缓冲区大小限制
// 2. 操作系统限制
// 3. MTU限制
// 4. 网络拥塞
func SendData(conn *net.TCPConn, data []byte) error {
	totalLen := len(data)
	writeLen := 0
	for {
		len, err := conn.Write(data[writeLen:])
		if err != nil {
			return err
		}
		writeLen = writeLen + len
		if writeLen >= totalLen {
			break
		}
	}
	return nil
}
