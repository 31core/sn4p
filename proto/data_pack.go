package proto

import (
	"time";
	"crypto/sha256"
)

type DataPack struct {
	Content DataPackContent
	Security DataPackSecurity
}

type DataPackContent struct {
	Header 		[20]byte
	TimeStamp 	int64
	sha256 		[32]byte
	Data 		[]byte
}

type DataPackSecurity struct {
	TimeStamp	int64 //第一次连接时间戳
	AES128		[16]byte
}

/* 组建数据包 */
func (self *DataPack) Build() []byte {
	var data []byte

	self.Content.TimeStamp = time.Now().Unix()
	self.Content.sha256 = sha256.Sum256(self.Content.Data)

	data = append(data, self.Content.Header[:]...)
	for i := 8; i > 0; i-- {
		data = append(data, byte(self.Content.TimeStamp >> (8 * i)))
	}

	data = append(data, self.Content.sha256[:]...)
	data = append(data, self.Content.Data...)
	return data
}

func (self *DataPack) Parse(data []byte) {
	copy(self.Content.Header[:], data[0:len(self.Content.Header)])
	copy(self.Content.sha256[:], data[28:28 + len(self.Content.sha256)])
	self.Content.Data = data[60:]
}
