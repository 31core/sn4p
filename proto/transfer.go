package proto

import (
	"net"
)

var conn net.Conn

type Transfer struct {
	TargetIP   string
	TargetPort uint16
	TimeStamp  int64 //第一次连接时间戳
	AES128     [16]byte
	DataPack   *DataPack
}

/* 发送数据 */
func (self Transfer) Send() error {
	data := self.DataPack.Build()
	data = EncryptAES(data, self.AES128)
	_, err := conn.Write(data)
	return err
}

/* 接收数据 */
func (self Transfer) Receive() error {
	data := make([]byte, 1024)
	size, err := conn.Read(data)
	if err != nil {
		return err
	}
	data = data[:size]
	data = DecryptAES(data, self.AES128)
	self.DataPack.Parse(data)
	return nil
}

/* 关闭连接 */
func (self Transfer) Close() error {
	return conn.Close()
}
