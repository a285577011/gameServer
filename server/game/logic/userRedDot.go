package logic

import (
	"github.com/name5566/leaf/gate"
	"server/game/server"
	"server/msg"
)

type userRedDot struct {
	Base
}

func NewUserRedDot() *userRedDot {
	obj := &userRedDot{}
	return obj
}
func (this *userRedDot) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*msg.UserRedDot)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info:=make(map[string]interface{})
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return


}
