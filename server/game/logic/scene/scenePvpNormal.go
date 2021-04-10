package scene

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
	"server/game/server"
	"time"
)

type ScenePvpNormal struct {
	SceneId   string
	IncrId    int64
	round     int32
	Persons   map[int64]*PvpRole
	Monsters  map[int64]*PvpRole
	Roles     map[int64]*PvpRole
	RoundLogs []interface{}
	secTime   int64
	fightChan chan int64
}

func NewScenePvpNormal() *ScenePvpNormal {
	obj := &ScenePvpNormal{}
	obj.IncrId = 0
	obj.round = 0
	obj.Roles = make(map[int64]*PvpRole)
	obj.Persons = make(map[int64]*PvpRole)
	obj.Monsters = make(map[int64]*PvpRole)
	obj.fightChan = make(chan int64)
	obj.createFight() //战斗逻辑监听就绪
	return obj
}
func (this *ScenePvpNormal) createFight() {
	// crn expr
	go func() { //起一个战斗队列
		for roleId := range this.fightChan {
			log.Debug("fight", roleId)
			Role := this.Roles[roleId]
			if Role == nil || len(Role.Target) == 0 {
				Role.closeFightChan <- 1 //关闭个人自动战斗
				continue
			}
			this.DoRoleFight(Role)
		}
	}()
}

func (this *ScenePvpNormal) SelectTarget(role *PvpRole, beAtkRoleId int64) {
	if len(role.Target) > 0 && beAtkRoleId == 0 {
		role.closeFightChan <- 1 //关闭个人自动战斗
		role.Target = map[int64]int64{}
		return
	}
	if len(role.Target) == 0 { //无目标到柚木不爱时丢一个战斗
		this.fightChan <- role.Id
		go func() { //并且定时自行自动打
			d := timer.NewDispatcher(10)
			// crn expr
			cronExpr, err := timer.NewCronExpr("* * * * * *")
			if err != nil {
				return
			}
			// cron
			var c *timer.Cron
			c = d.CronFunc(cronExpr, func() {
				fmt.Println("My name is Leaf")
				select {
				case <-role.closeFightChan:
					log.Debug("fight end")
					c.Stop()
				default:
					this.fightChan <- role.Id
					log.Debug("fight next")
				}
			})
			for c := range d.ChanTimer {
				c.Cb()
			}
		}()
	}
	role.Target[beAtkRoleId] = beAtkRoleId

}
func (this *ScenePvpNormal) GetSceneId() string {
	return this.SceneId
}
func (this *ScenePvpNormal) InitRoles() {
}
func (this *ScenePvpNormal) Reset() {
}
func (this *ScenePvpNormal) DoFightAfterSetTarget() {
}
func (this *ScenePvpNormal) DoFight() {
}
func (this *ScenePvpNormal) FindMonsterTarget() {
}
func (this *ScenePvpNormal) DoRoleHurt() {
}
func (this *ScenePvpNormal) PushFightInfo() {
}
func (this *ScenePvpNormal) ClearRoundFightInfo() {
}
func (this *ScenePvpNormal) PersonJoin(person *PvpRole) {
	if this.CheckExist(person.Uid) != nil {
		return
	}
	uIds := this.GetSceneUids()
	person.Id = this.getRoleIncrId()
	this.Persons[person.Uid] = person
	this.Roles[person.Id] = person
	this.AddRoles(person)
	message := make(map[string]interface{})
	message["uId"] = person.Uid
	message["action"] = PERSON_ACTION_JOIN
	message["sceneId"] = this.SceneId
	message["onNum"] = len(uIds)
	message["role"] = person
	message["tNum"] = 1
	server.BroadCastByUidS(uIds, message, TOFRONT_PVP_PERSON)
}
func (this *ScenePvpNormal) GetSceneUids() []int64 {
	res := []int64{}
	if len(this.Persons) == 0 {
		return res
	}
	for _, v := range this.Persons {
		res = append(res, v.Uid)
	}
	return res
}
func (this *ScenePvpNormal) GetMonster() map[int64]*PvpRole {
	return this.Monsters
}
func (this *ScenePvpNormal) GetPerson() map[int64]*PvpRole {
	return this.Persons
}
func (this *ScenePvpNormal) GetRoles() map[int64]*PvpRole {
	return this.Roles
}
func (this *ScenePvpNormal) PersonLeave(person *PvpRole) {
	if this.CheckExist(person.Uid) != nil {
		return
	}
	this.DelRoles(person)
	delete(this.Persons, person.Uid)
	uIds := this.GetSceneUids()
	message := make(map[string]interface{})
	message["uId"] = person.Uid
	message["action"] = PERSON_ACTION_LEAVE
	message["sceneId"] = this.SceneId
	message["onNum"] = len(uIds)
	message["role"] = person
	server.BroadCastByUidS(uIds, message, TOFRONT_PVP_PERSON)
}
func (this *ScenePvpNormal) AddMonster(monster *PvpRole) {
	monster.Id = this.getRoleIncrId()
	this.AddRoles(monster)
	this.Monsters[monster.Id] = monster
}
func (this *ScenePvpNormal) MonsterLeave(monster *PvpRole) {
	this.DelRoles(monster)
	delete(this.Monsters, monster.Id)
}
func (this *ScenePvpNormal) MonsterNextTime(monster *PvpRole) {
	if monster.NextTime != 0 {
		return
	}
	monster.NextTime = time.Now().Unix() + 5
	uIds := this.GetSceneUids()
	message := make(map[string]interface{})
	message["id"] = monster.Id
	//
	message["nextTime"] = monster.NextTime
	message["sceneId"] = this.SceneId
	message["hp"] = monster.Hp
	message["maxHp"] = monster.Maxhp
	message["cId"] = monster.Cid
	message["isBuff"] = 0
	message["type"] = monster.Type
	server.BroadCastByUidS(uIds, message, TOFRONT_PVP_BOSS_ATTR)
}
func (this *ScenePvpNormal) AddRoles(role *PvpRole) {
	this.Roles[role.Id] = role
}
func (this *ScenePvpNormal) DelRoles(role *PvpRole) {
	delete(this.Roles, role.Id)
}
func (this *ScenePvpNormal) CheckExist(uId int64) *PvpRole {
	role := this.Persons[uId]
	return role
}
func (this *ScenePvpNormal) getRoleIncrId() int64 {
	this.IncrId++
	return this.IncrId
}
func (this *ScenePvpNormal) DoRoleFight(atkRole *PvpRole) bool {
	if atkRole.Target == nil || len(atkRole.Target) == 0 {
		return false
	}
	logItems := []map[string]interface{}{}
	for _, tarGet := range atkRole.Target {
		tarGetRole := this.Roles[tarGet]
		if tarGetRole == nil {
			continue
		}
		if tarGetRole.Hp <= 0 {
			continue
		}
		var hurt float64 = 100
		tarGetRole.Hp -= hurt
		logItem := make(map[string]interface{})
		logItem["t"] = LOG_TYPE_NORMAL
		logItem["fId"] = atkRole.Id
		logItem["tId"] = tarGetRole.Id
		logItem["ht"] = hurt
		logItem["rht"] = hurt
		logItem["hp"] = tarGetRole.Hp
		logItem["ct"] = 0
		cht := make(map[string]interface{})
		cht["ht"] = hurt
		cht["rht"] = hurt
		logItem["cht"] = cht
		logItem["miss"] = 0
		logItem["isTask"] = 0
		logItems = append(logItems, logItem)
		if tarGetRole.Hp <= 0 {
			delete(atkRole.Target, tarGetRole.Id)
		}
	}
	uIds := this.GetSceneUids()
	pushMsg := make(map[string]interface{})
	logs := make(map[int]interface{})
	formatLog := this.formatLogItems(logItems)
	logs[0] = formatLog
	pushMsg["logs"] = logs
	server.BroadCastByUidS(uIds, pushMsg, TOFRONT_PVP_FIGHT)
	if len(atkRole.Target) > 0 {
		return true
	}
	return false
}
func (this *ScenePvpNormal) formatLogItems(logItems []map[string]interface{}) map[string]interface{} {
	logItem := make(map[string]interface{})
	if len(logItems) == 1 {
		logItem = logItems[0]
	} else {
		fId := logItems[0]["fId"]
		logItem = map[string]interface{}{"t": LOG_TYPE_SKILL, "fId": fId}
		tmpLogs := []interface{}{}
		for _, v := range logItems {
			tmpLogs = append(tmpLogs, v)
		}
		logItem["logs"] = tmpLogs
	}
	return logItem
}
