package main

import (
	"sn4p/proto"
)

func main() {
	var t proto.Transfer
	pack := proto.NewPack()
	t.DataPack = &pack
	t.Connect("127.0.0.1", 2233)
	t.Send()
}
