package main

import (
	"sn4p/proto"
)

func main() {
	var t proto.Transfer
	var pack proto.DataPack
	t.Target_ip = "127.0.0.1"
	t.Data_pack = &pack
	t.Send()
}