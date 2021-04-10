package gate

import (
	"server/game"
	"server/msg"
	_map "server/msg/map"
)

func init() {
	msg.Processor.SetRouter(&msg.Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.User{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Activity{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupWar{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Reconnet{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserRedDot{}, game.ChanRPC)
	msg.Processor.SetRouter(&_map.MapList{}, game.ChanRPC)
	msg.Processor.SetRouter(&_map.MapJoin{}, game.ChanRPC)
	msg.Processor.SetRouter(&_map.MapSelectTarget{}, game.ChanRPC)
	msg.Processor.SetRouter(&_map.MapLeave{}, game.ChanRPC)
}
