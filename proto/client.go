package proto

import (
	"net"
	"fmt"
)

type Transfer struct {
	TargetIP string
	DataPack *DataPack
}

var conn net.Conn

func (self Transfer) Conect() error {
	var err error
	conn, err = net.Dial("tcp", self.TargetIP + ":2233")
	if err != nil {
		fmt.Printf("error: unable to connect: %s\n", self.TargetIP)
	}
	return err
}

func (self Transfer) Send() error {
	_, err := conn.Write(self.DataPack.Build())
	return err
}

func (self Transfer) Close() error {
	return conn.Close()
}