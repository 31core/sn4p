package main

import (
	"sn4p/proto"
)

func main() {
	for {
		var t proto.Transfer
		pack := proto.NewPack()
		t.DataPack = &pack
		t.Accept()
		for {
		}
	}
}
