package scene

type ScenePvp interface {
	InitRoles()
	DoRoleFight(atkRole *PvpRole) bool
	DoFightAfterSetTarget()
	Reset()
	DoFight()
	DoRoleHurt()
	PushFightInfo()
	FindMonsterTarget()
	ClearRoundFightInfo()
	PersonJoin(person *PvpRole)
	AddMonster(monster *PvpRole)
	PersonLeave(person *PvpRole)
	CheckExist(uId int64) *PvpRole
	GetSceneId() string
	GetSceneUids() []int64
	GetMonster() map[int64]*PvpRole
	GetPerson() map[int64]*PvpRole
	GetRoles() map[int64]*PvpRole
	SelectTarget(role *PvpRole,beAtkRoleId int64)
}
