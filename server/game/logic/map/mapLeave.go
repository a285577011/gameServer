package _map

import (
	"github.com/name5566/leaf/gate"
	"server/game/logic"
	"server/game/logic/scene"
	"server/game/server"
	_map "server/msg/map"
	"strings"
)

type mapLeave struct {
	logic.Base
}

func NewMapLeave() *mapLeave {
	obj := &mapLeave{}
	return obj
}
func (this *mapLeave) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*_map.MapLeave)
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
	sceneObj.PersonLeave(userRole)
	info := make(map[string]interface{})

	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return
}
func (this *mapLeave) OffLine(agent gate.Agent) {
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		return
	}
	gameUser := server.GetUser(uId)
	stageSlice := strings.Split(gameUser.Stage, "_")
	if len(stageSlice) == 0 {
		return
	}
	cId := stageSlice[1]
	sceneObj := scene.NewScene().GetOrCreateScenePvp(cId)
	userRole := sceneObj.CheckExist(uId)
	if userRole == nil {
		return
	}
	sceneObj.PersonLeave(userRole)
}
