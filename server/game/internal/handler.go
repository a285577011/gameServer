package internal

import (
	"reflect"
	"server/game/logic"
	lmap "server/game/logic/map"
	"server/msg"
	_map "server/msg/map"
)

func init() {
	handler(&msg.Login{},logic.NewUserLogin().Handle)
	handler(&msg.User{},logic.NewUserGetInfo().Handle)
	handler(&msg.Activity{},logic.NewActivityGetInfo().Handle)
	handler(&msg.GroupWar{},logic.NewGroupWarInfo().Handle)
	handler(&msg.Reconnet{},logic.NewUserReconn().Handle)
	handler(&msg.UserRedDot{},logic.NewUserRedDot().Handle)
	handler(&_map.MapList{},lmap.NewMapList().Handle)
	handler(&_map.MapJoin{},lmap.NewMapJoin().Handle)
	handler(&_map.MapSelectTarget{},lmap.NewMapSelectTarget().Handle)
	handler(&_map.MapLeave{},lmap.NewMapLeave().Handle)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
