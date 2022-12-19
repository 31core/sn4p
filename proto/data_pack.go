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
	Size          uint64 //content size
	Data          []byte
}

func NewPack() DataPack {
	var data_pack DataPack
	data_pack.Version = ProtoVersion
	data_pack.ClientVersion = [3]byte{ClientVersionMajor, ClientVersionMinimal, ClientVersionPatch}
	return data_pack
}

/* build datapack */
func (d *DataPack) Build() []byte {
	var data []byte
	buffer := bytes.NewBuffer([]byte{})

	d.TimeStamp = uint64(time.Now().Unix())
	d.Sha256 = sha256.Sum256(d.Data)
	d.Size = uint64(len(d.Data))

	data = append(data, d.Version)
	data = append(data, d.ClientVersion[:]...)
	data = append(data, d.Type)

	binary.Write(buffer, binary.BigEndian, d.TimeStamp)
	data = append(data, buffer.Bytes()...)

	data = append(data, d.Sha256[:]...)

	buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, d.Size)
	data = append(data, buffer.Bytes()...)

	data = append(data, d.Data...)
	return data
}

func (d *DataPack) Parse(data []byte) {
	d.Version = data[0]
	d.ClientVersion = [3]byte{data[1], data[2], data[3]}
	d.Type = data[4]

	buffer := bytes.NewBuffer(data[5:13])
	binary.Read(buffer, binary.BigEndian, d.TimeStamp)

	copy(d.Sha256[:], data[13:45])

	buffer = bytes.NewBuffer(data[45:53])
	binary.Read(buffer, binary.BigEndian, d.Size)

	copy(d.Data[:], data[53:])
}
