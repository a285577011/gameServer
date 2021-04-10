package logic

import (
	"github.com/name5566/leaf/gate"
	"server/game/server"
	"server/msg"
)

type userReconn struct {
	Base
}

func NewUserReconn() *userReconn {
	obj := &userReconn{}
	return obj
}
func (this *userReconn) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*msg.Reconnet)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info:=make(map[string]interface{})
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return


}
