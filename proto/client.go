package proto

import (
	"net"
	"time"
	"fmt"
)

type Transfer struct {
	TargetIP string
	DataPack *DataPack
}

var conn net.Conn

func (self *Transfer) Conect() error {
	var err error
	conn, err = net.Dial("tcp", self.TargetIP + ":2233")
	if err != nil {
		fmt.Printf("error: unable to connect: %s\n", self.TargetIP)
	}
	self.DataPack.Security.TimeStamp = time.Now().Unix()
	return err
}
