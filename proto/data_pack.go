package proto

import (
	"time";
	"crypto/sha256"
)

type DataPack struct {
	Header 		[20]byte
	TimeStamp 	int64
	sha256 		[32]byte
	Data 		[]byte
}

/* 组建数据包 */
func (self *DataPack) Build() []byte {
	var data []byte

	self.TimeStamp = time.Now().Unix()
	self.sha256 = sha256.Sum256(self.Data)

	data = append(data, self.Header[:]...)
	for i := 8; i > 0; i-- {
		data = append(data, byte(self.TimeStamp >> (8 * i)))
	}

	data = append(data, self.sha256[:]...)
	data = append(data, self.Data...)
	return data
}

func (self *DataPack) Parse(data []byte) {
	copy(self.Header[:], data[0:len(self.Header)])
	copy(self.sha256[:], data[28:28 + len(self.sha256)])
	self.Data = data[60:]
}
