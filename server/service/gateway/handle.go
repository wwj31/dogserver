package gateway

// 处理其他服务向gateway发送的消息
func (s *GateWay) InnerHandler(sourceId string, v interface{}) {
	switch v.(type) {
	default:
	}
}
