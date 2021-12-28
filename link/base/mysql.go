package base

import (
	"errors"
	"fmt"
	"github.com/IEatLemons/BaseGin/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"time"
)

var MysqlGroup *xorm.EngineGroup

func InitEngineGroup(conf ...*config.MySQL) (err error) {
	if len(conf) <= 0 {
		err = errors.New("config is nil ")
		return
	}
	var MasterConf config.MySQL
	var cfg []string
	for k, v := range conf {
		if k == 0 {
			MasterConf = *v
		}
		mysqlLink := fmt.Sprintf("%v:%v@(%v:%d)/%v?charset=%v", v.Username, v.Password, v.Host, v.Port, v.Db, v.Charset)
		cfg = append(cfg, mysqlLink)
	}

	db, err := xorm.NewEngineGroup("mysql", cfg)
	if err != nil {
		return
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(MasterConf.MaxIdleConns)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(MasterConf.MaxOpenConns)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	db.ShowSQL(true)
	MysqlGroup = db
	return
}
