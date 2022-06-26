package proto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type DataPack struct {
	Version       byte
	ClientVersion [3]byte
	Type          byte
	TimeStamp     uint64
	Sha256        [32]byte
	Size          uint64
	Data          []byte
}

func NewPack() DataPack {
	var data_pack DataPack
	data_pack.Version = ProtoVersion
	data_pack.ClientVersion = [3]byte{ClientVersionMajor, ClientVersionMinimal, ClientVersionPatch}
	return data_pack
}

/* 组建数据包 */
func (self *DataPack) Build() []byte {
	var data []byte
	buffer := bytes.NewBuffer([]byte{})

	self.TimeStamp = uint64(time.Now().Unix())
	self.Sha256 = sha256.Sum256(self.Data)
	self.Size = uint64(len(self.Data))

	data = append(data, self.Version)
	data = append(data, self.ClientVersion[:]...)
	data = append(data, self.Type)

	binary.Write(buffer, binary.BigEndian, self.TimeStamp)
	data = append(data, buffer.Bytes()...)

	data = append(data, self.Sha256[:]...)

	buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, self.Size)
	data = append(data, buffer.Bytes()...)

	data = append(data, self.Data...)
	return data
}

func (self *DataPack) Parse(data []byte) {
	self.Version = data[0]
	self.ClientVersion = [3]byte{data[1], data[2], data[3]}
	self.Type = data[4]

	buffer := bytes.NewBuffer(data[5:13])
	binary.Read(buffer, binary.BigEndian, self.TimeStamp)

	copy(self.Sha256[:], data[13:45])

	buffer = bytes.NewBuffer(data[45:53])
	binary.Read(buffer, binary.BigEndian, self.Size)

	copy(self.Data[:], data[53:])
}
