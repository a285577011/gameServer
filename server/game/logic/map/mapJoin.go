package _map

import (
	"github.com/name5566/leaf/gate"
	"server/game/logic"
	"server/game/logic/scene"
	"server/game/server"
	_map "server/msg/map"
	"strconv"
)

type mapJoin struct {
	logic.Base
}

func NewMapJoin() *mapJoin {
	obj := &mapJoin{}
	return obj
}
func (this *mapJoin) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*_map.MapJoin)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info := make(map[string]interface{})
	cIdStr := strconv.Itoa(int(m.CId))
	sceneObj := scene.NewScene().GetOrCreateScenePvp(cIdStr)
	userRole := sceneObj.CheckExist(uId)
	if userRole == nil {
		userRedisInfo := server.GetUserRedisInfo(strconv.Itoa(int(uId)))
		userRole = scene.NewPvpRole().CreatePerson(userRedisInfo)
	}
	sceneObj.PersonJoin(userRole)
	gameUser := server.GetUser(uId)

	gameUser.Stage = "main_" + strconv.Itoa(int(m.CId))
	info["id"] = userRole.Id
	info["hp"] = userRole.Hp
	info["maxHp"] = userRole.Maxhp
	info["scene"] = sceneObj.GetSceneId()
	info["onNum"] = len(sceneObj.GetSceneUids())
	ms := []interface{}{}
	msMap := sceneObj.GetMonster()
	for _, v := range msMap {
		ms = append(ms, v)
	}
	ps := []interface{}{}
	psMap := sceneObj.GetPerson()
	for _, v := range psMap {
		ps = append(ps, v)
	}
	info["ms"] = ms
	info["ps"] = ps
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return
}
