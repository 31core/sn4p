package proto

import (
	"crypto/rand"
	"net"
	"time"
)

/* 等待连接 */
func (self *Transfer) Accept() error {
	server, err := net.Listen("tcp", ":2233")
	conn, err = server.Accept()
	self.TimeStamp = time.Now().Unix()

	data := make([]byte, 1024)
	/* 从客户端接收公钥 */
	conn.Read(data)
	publickey := LoadPublicKey(data)

	/* 向客户端发送AES密钥 */
	rand.Read(self.AES128[:])
	conn.Write(EncryptRSA(self.AES128[:], publickey))
	return err
}
