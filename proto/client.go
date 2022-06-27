package proto

import (
	"fmt"
	"net"
	"time"
)

func (self *Transfer) Connect(ip string, port uint16) error {
	var err error
	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", self.TargetIP, self.TargetPort))
	if err != nil {
		fmt.Printf("error: unable to connect: %s\n", self.TargetIP)
		return err
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
	copy(self.AES128[:], DecryptRSA(data, privatekey))
	return nil
}
