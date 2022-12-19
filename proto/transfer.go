package proto

import (
	"net"
)

var conn net.Conn

type Transfer struct {
	TargetIP   string
	TargetPort uint16
	TimeStamp  int64 //time stamp of the first connection
	AES128     [16]byte
	DataPack   *DataPack
}

/* send data */
func (t Transfer) Send() error {
	data := t.DataPack.Build()
	data = EncryptAES(data, t.AES128)
	_, err := conn.Write(data)
	return err
}

/* receive data */
func (t Transfer) Receive() error {
	data := make([]byte, 1024)
	size, err := conn.Read(data)
	if err != nil {
		return err
	}
	data = data[:size]
	data = DecryptAES(data, t.AES128)
	t.DataPack.Parse(data)
	return nil
}

/* close connection */
func (t Transfer) Close() error {
	return conn.Close()
}
