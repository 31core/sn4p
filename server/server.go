package main

import (
	"net"
	"sn4p/proto"
)

func main() {
	for {
		server, _ := net.Listen("tcp", ":2233")
		for {
			var data_pack proto.DataPack
			data := make([]byte, 1024)
			con, _ := server.Accept()
			con.Read(data)
			data_pack.Parse(data)
		}
	}
}
