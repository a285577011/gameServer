package _map

import (
	"github.com/name5566/leaf/gate"
	"server/conf/game"
	"server/game/logic"
	"server/game/server"
	_map "server/msg/map"
	"strconv"
	"strings"
)

type mapList struct {
	logic.Base
}

func NewMapList() *mapList {
	obj := &mapList{}
	return obj
}
func (this *mapList) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*_map.MapList)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info := make(map[string]interface{})
	missConf := game.GetConf("mission")
	monsterConf := game.GetConf("monster")
	lcd := 0
	mapInfo := make(map[int]interface{})
	for _, mId := range m.MId {
		mConf := missConf[strconv.Itoa(mId)].(map[string]interface{})
		mapItem := make(map[string]interface{})
		mapItem["cd"] = lcd
		bossIds := strings.Split(mConf["bossId"].(string), ",")
		msItems := []interface{}{}
		for _, bossId := range bossIds {
			msConf := monsterConf[bossId].(map[string]interface{})
			msItem := make(map[string]interface{})
			msItem["hp"] = msConf["hp"]
			msItem["maxhp"] = msConf["hp"]
			msItem["cId"] = msConf["id"]
			msItem["nextTime"] = 0
			msItems = append(msItems, msItem)
		}
		mapItem["ms"] = msItems
		mapInfo[mId] = mapItem
	}
	info["l"] = mapInfo
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return
	// 给发送者回应一个 Hello 消息
}
