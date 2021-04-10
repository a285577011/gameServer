package _map

import "server/msg"

func init() {
	msg.Processor.Register(&MapList{}, "m.e")
}

type MapList struct {
	msg.Common
	MId []int
}
