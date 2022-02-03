package main

import (
	"sn4p/proto"
)

func main() {
	var t proto.Transfer
	var pack proto.DataPack
	t.TargetIP = "127.0.0.1"
	t.DataPack = &pack
	t.Conect()
	t.Send()
}