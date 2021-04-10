package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"server/db/mysql"
	"server/db/redis"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
	mysql.MysqlCenterDb.Close()
	mysql.MysqlServerDb.Close()
	redis.RedisDb.Close()
}
