package base

import (
	"github.com/go-redis/redis"
	"github.com/xormplus/xorm"
)

type Link struct {
	Mysql *xorm.EngineGroup
	Redis *redis.Client
}

func NewLink() *Link {
	return &Link{
		Mysql: MysqlGroup,
		Redis: Redis,
	}
}
