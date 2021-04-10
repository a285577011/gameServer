package logic

import (
	"github.com/name5566/leaf/gate"
	"server/game/server"
	"server/msg"
)

type groupWarInfo struct {
	Base
}

func NewGroupWarInfo() *groupWarInfo {
	obj := &groupWarInfo{}
	return obj
}
func (this *groupWarInfo) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*msg.GroupWar)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info := make(map[string]interface{})
	idata := make(map[string]interface{})
	groupData := make(map[string]interface{})
	groupData["cId"] = 1
	groupData["lcd"] = 0
	groupData["ms"] = []interface{}{}
	groupData["win"] = 0
	groupData["team"] = 0
	groupData["st"] = 0
	groupData["ret"] = 0
	groupData["et"] = 0
	groupData["dt"] = 0
	groupData["status"] = 0
	idata["groupWar"] = groupData
	info["i"] = idata
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return

}
