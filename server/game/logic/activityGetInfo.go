package logic

import (
	"encoding/json"
	"github.com/name5566/leaf/gate"
	"server/game/server"
	"server/msg"
)

type activityGetInfo struct {
	Base
}

func NewActivityGetInfo() *activityGetInfo {
	obj := &activityGetInfo{}
	return obj
}
func (this *activityGetInfo) Handle(args []interface{}) {
	agent := args[1].(gate.Agent)
	m := args[0].(*msg.Activity)
	uId := server.GetUserByAgent(agent)
	if uId == 0 {
		this.ErrorMsg(m.Route, m.ReqId, nil, "非法用户", 404)
	}
	info := make(map[string]interface{})
	info["l"] = []string{}
	info["dl"] = make(map[string]interface{})
	info["v"] = 20
	saltJson := `{"activity":{"sevenDayLogin":[{"idx":3,"salt":"5146679bcbda0b3e63db8991752afc32"}],"sign":[{"idx":4,"salt":"ae5d5d7d534aaff1b778d997b3004b3b"}],"dailyRecharge":[{"idx":0,"salt":""}],"totalRecharge":[{"idx":0,"salt":""}],"rechargeDay":[{"idx":0,"salt":""}],"rechargeMonth":[{"idx":0,"salt":""}]},"sevenday_task":{"sevenDayTask":[{"idx":2,"salt":"7fd0101d370d45b26151177ed69f7c6a"}]},"activity_gift":{"weeklyGift":[{"idx":0,"salt":""}],"monthlyGift":[{"idx":0,"salt":""}],"dailyGift":[{"idx":0,"salt":""}]},"activity_rank":{"battleRank":[{"idx":0,"salt":""}],"cardFpRank":[{"idx":0,"salt":""}],"levelRank":[{"idx":0,"salt":""}],"shenbingRank":[{"idx":0,"salt":""}],"equipFpRank":[{"idx":0,"salt":""}],"fpRank":[{"idx":0,"salt":""}],"guildFpRank":[{"idx":0,"salt":""}],"fateFpRank":[{"idx":0,"salt":""}],"teachFpRank":[{"idx":0,"salt":""}],"bootyFpRank":[{"idx":0,"salt":""}],"gcRank":[{"idx":0,"salt":""}]},"knight_order":{"knightOrder":[{"idx":0,"salt":""}]},"gift_shop":{"giftShop":[{"idx":0,"salt":""}]},"free_buy":{"freeBuy":[{"idx":0,"salt":""}]},"visit_capital":{"palace":[{"idx":0,"salt":""}]}}`
	salt := make(map[string]interface{})
	json.Unmarshal([]byte(saltJson), &salt)
	info["dl"] = salt
	res := this.Success(m.Route, m.ReqId, info)
	agent.WriteMsg(res)
	return

}
