package proto

import (
	"fmt"
	"net"
	"time"
)

func (self *Transfer) Conect() error {
	var err error
	conn, err = net.Dial("tcp", self.TargetIP+":2233")
	if err != nil {
		fmt.Printf("error: unable to connect: %s\n", self.TargetIP)
	}
	self.TimeStamp = time.Now().Unix()

	data := make([]byte, 1024)
	/* 向服务器发送公钥 */
	privatekey, publickey := GenerateRSAKey()
	b_publickey := DumpPublicKey(publickey)
	conn.Write(b_publickey)

	data = make([]byte, 1024)
	/* 从服务器接收AES密钥 */
	i, _ := conn.Read(data)
	data = data[:i]
	self.AES128 = DecryptRSA(data, privatekey)
	return err
}
