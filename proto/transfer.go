package proto

/* 发送数据 */
func (self Transfer) Send() error {
	_, err := conn.Write(self.DataPack.Build())
	return err
}

/* 接收数据 */
func (self Transfer) Receive() error {
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	self.DataPack.Parse(data)
	return err
}

/* 关闭连接 */
func (self Transfer) Close() error {
	return conn.Close()
}
