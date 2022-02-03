package proto

import (
	"net"
	"time"
)

/* 等待连接 */
func (self *Transfer) Accept() error {
	server, err := net.Listen("tcp", ":2233")
	conn, err = server.Accept()
	self.DataPack.Security.TimeStamp = time.Now().Unix()
	return err
}