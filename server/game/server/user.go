package server

import (
	"github.com/name5566/leaf/gate"
	"server/db/redis"
	"sync"
	"time"
)

var userMap map[int64]*User
var agentUserMap map[gate.Agent]int64

type User struct {
	Id           int64
	UserName     string
	Agent        gate.Agent
	OpenId       string
	SyncTime     int32
	SyncNum      int32
	LoginTime    int32
	Info         map[string]interface{}
	Stage        string
	ActivityTime int64
}
var rmu sync.RWMutex
func init() {
	userMap = make(map[int64]*User)
	agentUserMap = make(map[gate.Agent]int64)
}
func InitUser(uId int64, agent gate.Agent, openId string) *User {
	now := time.Now().Unix()
	rmu.RLock()
	userMap[uId] = &User{uId, "", agent, openId, 0, 0, 0, nil, "", now}
	rmu.RUnlock()
	agentUserMap[agent] = uId
	return userMap[uId]
}
func DelUser(agent gate.Agent) {
	uId := agentUserMap[agent]
	if uId == 0 {
		return
	}
	delete(agentUserMap, agent)
	rmu.RLock()
	delete(userMap, uId)
	rmu.RUnlock()
}
func GetUser(uId int64) *User {
	rmu.RLock()
	defer rmu.RUnlock()
	user, ok := userMap[uId]
	if ok {
		return user
	}
	return nil
}
func GetUserByAgent(agent gate.Agent) int64 {
	uId, ok := agentUserMap[agent]
	if ok {
		return uId
	}
	return 0
}
func GetUserRedisInfoByFields(uId string, field ...string) interface{} {
	redisKey := uId + "_user"
	if len(field) == 1 {
		rs := redis.RedisDb.HMGet(redisKey, field[0]).Val()
		return rs[0]
	}
	rs := redis.RedisDb.HMGet(redisKey, field...).Val()
	return rs
}
func GetUserRedisInfo(uId string) map[string]string {
	redisKey := uId + "_user"
	rs := redis.RedisDb.HGetAll(redisKey).Val()
	return rs
}
