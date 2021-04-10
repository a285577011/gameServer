package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	lmap "server/game/logic/map"
	"server/game/server"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	log.Debug("connect ...")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	server.DelUser(a)
	lmap.NewMapLeave().OffLine(a)
	log.Debug("connect close")
}
