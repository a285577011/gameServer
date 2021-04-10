package scene

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"server/conf/game"
	"server/msg/gameMap"
	"strconv"
)

type PvpRole struct {
	Id             int64
	Uid            int64
	Cid            int32
	Lv             int32
	Type           int8
	Atk            int32
	Hp             float64
	Maxhp          float64
	Target         map[int64]int64
	NextTime       int64
	AtkTime        int64
	Owner          int64
	Attr           map[string]interface{}
	closeFightChan chan int8 //关闭自动战斗的通道
}

func NewPvpRole() *PvpRole {
	obj := &PvpRole{}
	return obj
}

/**
创建一个怪物对象
*/
func (this *PvpRole) CreateMonster(msId string, msType int8) *PvpRole {
	monsterConf := game.GetConf("monster")
	msConf := monsterConf[msId].(map[string]interface{})
	return this.setRoleByConf(msConf, msType)

}
func (this *PvpRole) setRoleByConf(msConf map[string]interface{}, msType int8) *PvpRole {
	this.Hp = msConf["hp"].(float64)
	this.Maxhp = this.Hp
	this.Uid = 0
	this.Cid = int32(msConf["id"].(float64))
	this.Lv = int32(msConf["lv"].(float64))
	this.Type = msType
	this.Atk = 10000
	this.Attr = make(map[string]interface{})
	this.Attr["atk"] = 1000
	this.Attr["def"] = 1000
	this.Attr["hp"] = 10000
	this.NextTime = 0
	this.Owner = 0
	this.Target = make(map[int64]int64)
	return this
}
func (this *PvpRole) CreatePerson(userInfo map[string]string) *PvpRole {
	uId, _ := strconv.ParseInt(userInfo["uId"], 10, 64)
	this.Uid = uId
	this.Cid = 0
	this.Atk = 10000
	var attrInfo map[string]interface{}
	json.Unmarshal([]byte(userInfo["attr"]), &attrInfo)
	attrInfo = gameMap.FormatAttrItem(attrInfo)
	log.Debug("userrole", attrInfo)
	this.Attr = make(map[string]interface{})
	this.Attr = attrInfo
	lv, _ := strconv.ParseInt(userInfo["uId"], 10, 32)
	this.Lv = int32(lv)
	this.Type = PVP_ROLE_PERSON
	this.Hp = this.Attr["hp"].(float64)
	this.Maxhp = this.Hp
	this.Target = make(map[int64]int64)
	return this
}

//public $hp = 0;                             // 当前血量
//    public $maxHp = 0;                          // 最大血量
//    public $initMaxHp = 0;                      // 初始血量
//    public $target = 0;                         // 选择的目标
//    public $sltId = 0;                          // 选择的目标顺序，用于先后手计算
//    public $nextTime = 0;                       // Boss下次存活时间
//    public $callId = 0;                         // 召唤红名捕快的ID
//    public $atkType = PVP::ROLE_ATK_DO;         // 1:主动攻击;2:被动反击
//    public $extHp = 0;                          // 额外血量
//    public $atkTime = 0;                        // 上次攻击的时间(秒)，等于1时表示选中了目标，还没有攻击过
//    public $atkTimePhase = 0;                   // 攻击的时间段值
//    public $atkGrap = PVP::ATK_GAP_TIME;        // 攻击的时间间隔
//    public $belongId = 0;                       // 主线怪归属者uId
//    public $auto = 0;                           // 自动战斗(挂机打怪)
//    public $isNew = 2;                          // 是否是新手 1:是;2:否
//    public $owner = 0;                          // 怪物归属者
//    public $effect = [];                        // 特效
//    public $connect = 0;                        // 分组关联
//    public $callEndTime = 0;                    // 召唤的上阵生效截止时间（上阵npc）
//    public $status = 0;                         // 状态 具体场景定义
//    public $muiltAtk = 1;                         //多段攻击 默认1次
//    public $cardAc = 0;                         //武学大师
//    public $cardSkillVal = 0;                   //绝学气力值
//    public $cardSkillLv = 0;                   //绝学技能等级
