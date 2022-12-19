package proto

import (
	"crypto/rand"
	"net"
	"time"
)

/* wait for connection */
func (t *Transfer) Accept() error {
	server, err := net.Listen("tcp", ":2233")
	if err != nil {
		return err
	}
	conn, err = server.Accept()
	t.TimeStamp = time.Now().Unix()

	data := make([]byte, 1024)
	/* receive public key from client */
	conn.Read(data)
	publickey := LoadPublicKey(data)

	/* send AES key to client */
	rand.Read(t.AES128[:])
	conn.Write(EncryptRSA(t.AES128[:], publickey))
	return err
}
