package _map

import "server/msg"

func init() {
	msg.Processor.Register(&MapLeave{}, "m.b")
}

type MapLeave struct {
	msg.Common
	CId int
}
