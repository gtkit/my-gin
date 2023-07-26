package dao

import (
	"time"

	"github.com/gtkit/gormlog"
	"github.com/gtkit/orm"
	"gorm.io/gorm"

	"github.com/gtkit/logger"

	"ydsd_gin/config"
)

func initMysql() *gorm.DB {
	// 设置 mysql 链接信息
	orm.MysqlConfig(
		orm.Host(config.GetString("database.host")),
		orm.Port(config.GetString("database.port")),
		orm.DbType(config.GetString("database.dbtype")),
		orm.Name(config.GetString("database.name")),
		orm.User(config.GetString("database.username")),
		orm.WithPassword(config.GetString("database.password")),
	)
	// 配制 gorm
	orm.GormConfig(
		orm.PrepareStmt(true),
		orm.SkipDefaultTransaction(true),
		orm.GormLog(gormlog.L(logger.Zlog(), logger.Sql())), // 此处注意,logger日志需要先实例化
	)
	mydb := orm.NewMysql()

	sqlDB, err := mydb.DB()
	if err != nil {
		panic(err)
	}

	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.maxopenconn"))
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	sqlDB.SetMaxIdleConns(config.GetInt("database.maxidleconn"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.maxlifeseconds")) * time.Second)
	logger.Info("MySql connect success!")
	return mydb
}
