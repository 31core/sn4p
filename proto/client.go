package proto

import (
	"fmt"
	"net"
	"time"
)

func (t *Transfer) Connect(ip string, port uint16) error {
	var err error
	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", t.TargetIP, t.TargetPort))
	if err != nil {
		fmt.Printf("error: unable to connect: %s\n", t.TargetIP)
		return err
	}
	t.TimeStamp = time.Now().Unix()

	/* send public key to server */
	privatekey, publickey := GenerateRSAKey()
	b_publickey := DumpPublicKey(publickey)
	conn.Write(b_publickey)

	data := make([]byte, 1024)
	/* receive AES key from server */
	i, _ := conn.Read(data)
	data = data[:i]
	copy(t.AES128[:], DecryptRSA(data, privatekey))
	return nil
}
