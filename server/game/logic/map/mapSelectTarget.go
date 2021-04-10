package _map

import (
	"github.com/name5566/leaf/gate"
	"server/game/logic"
	"server/game/logic/scene"
	"server/game/server"
	_map "server/msg/map"
	"strings"
)

type mapSelectTarget struct {
	logic.Base
}

func NewMapSelectTarget() *mapSelectTarget {
	obj := &mapSelectTarget{}
	return obj
}
func (this *mapSelectTarget) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*_map.MapSelectTarget)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	gameUser := server.GetUser(uId)
	stageSlice := strings.Split(gameUser.Stage, "_")
	if len(stageSlice) == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "用户不在地图内", 1001)
	}
	cId := stageSlice[1]
	sceneObj := scene.NewScene().GetOrCreateScenePvp(cId)
	userRole := sceneObj.CheckExist(uId)
	if userRole == nil {
		this.ErrorMsg(m.Route, m.ReqId, nil, "用户不在地图内", 1002)
	}
	sceneObj.SelectTarget(userRole, m.Id)
	info := make(map[string]interface{})

	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return
}
