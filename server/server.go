package main

import (
	"sn4p/proto"
)

func main() {
	for {
		var t proto.Transfer
		var pack proto.DataPack
		t.DataPack = &pack
		t.Accept()
		for {}
	}
}
