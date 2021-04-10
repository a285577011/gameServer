package _map

import "server/msg"

func init() {
	msg.Processor.Register(&MapSelectTarget{}, "m.d")
}

type MapSelectTarget struct {
	msg.Common
	Id int64
}
