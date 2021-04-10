package logic

import (
	"github.com/name5566/leaf/gate"
	"server/conf"
	"server/db/redis"
	"server/game/common"
	"server/game/model"
	"server/game/server"
	"server/msg"
	"strconv"
	"time"
)

type userLogin struct {
}

func NewUserLogin() *userLogin {
	obj := &userLogin{}
	return obj
}
func (this *userLogin) Handle(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Login)
	hash := m.Hash
	skey := conf.SessionSignKey
	hashData := common.DecryptOpenssl(hash, skey)
	res := make(map[string]interface{})
	// 消息的发送者
	a := args[1].(gate.Agent)
	if len(hashData) == 0 {
		res = this.errorMsgLogin(m, hashData, "参数错误", 404)
		a.WriteMsg(res)
		return
	}
	openId, ok := hashData["openId"]
	if !ok {
		res = this.errorMsgLogin(m, hashData, "参数错误", 404)
		a.WriteMsg(res)
		return
	}

	if m.SId < 1 || openId == "" {
		res = this.errorMsgLogin(m, hashData, "参数错误", 404)
		a.WriteMsg(res)
		return
	}
	sql := "SELECT * from user where user_reg_id='" + openId.(string) + "'"
	userInfo, _ := model.ServerOrm.FindOne(sql)
	var uId int
	initRole := 0
	group := 0
	sex := 0
	if len(userInfo) == 0 {

	} else {
		uId, _ = strconv.Atoi(userInfo["user_id"])
		initRole, _ = strconv.Atoi(userInfo["user_init_role"])
		group, _ = strconv.Atoi(userInfo["user_group"])
		redisKey := userInfo["user_id"] + "_user"
		sex, _ = redis.RedisDb.HGet(redisKey, "sex").Int()
		now := strconv.FormatInt(time.Now().Unix(), 10)
		updateSql := "UPDATE user SET user_activity_time=" + now + " WHERE user_id=" + userInfo["user_id"]
		model.ServerOrm.Update(updateSql)
	}
	userGameData := server.GetUser(int64(uId))
	if userGameData != nil {
		userGameData.Agent.Close()
	}
	server.InitUser(int64(uId), a, openId.(string))
	info := make(map[string]interface{})
	info["uId"] = uId
	info["initRole"] = initRole
	info["group"] = group
	info["sex"] = sex
	res = this.successLogin(m, info)
	a.WriteMsg(res)
	return
	// 给发送者回应一个 Hello 消息
}
func (this *userLogin) errorMsgLogin(m *msg.Login, params map[string]interface{}, msg string, code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["msg"] = msg
	res["s"] = code
	res["t"] = "onAdd"
	res["mp"] = params
	res["reqId"] = m.ReqId
	return res
}
func (this *userLogin) successLogin(m *msg.Login, params map[string]interface{}) map[string]interface{} {
	params["s"] = 0
	params["t"] = "onAdd"
	params["reqId"] = m.ReqId
	return params
}
