package scene

const PVP_ROLE_PERSON int8 = 1
const PVP_ROLE_BOSS int8 = 2
const PVP_ROLE_MONSTER int8 = 3

// PVP相关推送
const TOFRONT_PVP_PERSON = 2001            // PVP有人员变动
const TOFRONT_PVP_FIGHT = 2002             // PVP打斗
const TOFRONT_PVP_DROP = 2003              // PVP结束掉落
const TOFRONT_PVP_TASK = 2004              // PVP结束更新任务进度
const TOFRONT_PVP_PERSON_ATTR = 2005       // PVP玩家属性变动
const TOFRONT_PVP_BOSS_ATTR = 2006         // PVPBoss属性变动
const TOFRONT_PVP_BOSS_REBIRTH = 2007      // BOSS复活
const TOFRONT_PVP_PLAYER_CANCLE_ATK = 2008 // 玩家取消攻击通知
const TOFRONT_PVP_HUNT = 2010              // PVP寻宝奖励通知
const TOFRONT_PVP_SCENE_CHANGE = 2011      // 场景信息变更通知
const TOFRONT_PVP_ATTACK_MSG = 2012        // 镖车攻击信息推送
const TOFRONT_PVP_SCENE_MSG = 2013         // 场景数据变更推送
const TOFRONT_PVP_SKILL_VAL = 2014         // PVP结束气力值变化
const TOFRONT_PVP_BOSS_CALL_NOTIFY = 2101  // boss复活通知

const PERSON_ACTION_JOIN = 1;           // 加入新人
const PERSON_ACTION_LEAVE = 2;          // 离开场景
const PERSON_ACTION_ROLE_CHANGE = 3;    // 场景角色信息变化（武学、神兵皮肤等）

const LOG_TYPE_NORMAL    = 1;        // 普通攻击
const LOG_TYPE_SKILL     = 2;        // 技能
const LOG_TYPE_BUFF      = 3;        // BUFF
const LOG_TYPE_NPC_ADDHP = 4;        // NPC回血