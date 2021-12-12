package proto

import (
	"net"
)

type Transfer struct {
	Target_ip string
	Data_pack *DataPack
}

func (self Transfer) Send() {
	conn, _ := net.Dial("tcp", self.Target_ip + ":2233")
	conn.Write(self.Data_pack.Build())
	conn.Close()
}