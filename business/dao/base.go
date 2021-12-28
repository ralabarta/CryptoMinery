package dao

import (
	"github.com/go-redis/redis"
	"github.com/xormplus/xorm"
)

type Condition interface {
	AddFindWhere(Session *xorm.Session)
}

type BaseDao struct {
	Mysql *xorm.EngineGroup
	Redis *redis.Client
}

func (D *BaseDao) SaveOne(Data interface{}, Id int, Session *xorm.Session) (affected int64, err error) {
	if Session == nil {
		Session = D.Mysql.NewSession()
	}
	if Id == 0 {
		affected, err = Session.InsertOne(Data)
	} else {
		affected, err = Session.ID(Id).Update(Data)
	}
	return
}
