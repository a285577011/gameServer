package _map

import "server/msg"

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	msg.Processor.Register(&MapJoin{}, "m.a")
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type MapJoin struct {
	msg.Common
	CId int64
}
