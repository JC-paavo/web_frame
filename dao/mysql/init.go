package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"web_frame/setting"
)

var DB *sqlx.DB

func Init(conf *setting.MysqlConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName, conf.OtherOptions) //user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Fatal("mysql connection error",
			zap.String("error", err.Error()),
		)
	}
	err = DB.Ping()
	if err != nil {
		zap.L().Fatal("mysql Ping error",
			zap.String("error", err.Error()),
		)
	}

	DB.SetMaxOpenConns(conf.MaxOpenConns)
	DB.SetMaxIdleConns(conf.MaxIdleConns)
	return
}
